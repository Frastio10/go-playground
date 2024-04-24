package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	addr string
	ln   net.Listener

	close chan struct{}
	msg   chan Message
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		close: make(chan struct{}),
		msg:   make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	defer ln.Close()

	s.ln = ln

	go s.AcceptLoop()
	<-s.close

	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Cannot accept", err)
			continue
		}

		fmt.Println("New client connected")

		go s.ReadPump(conn)
	}
}

func (s *Server) ReadPump(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Cannot read data: ", err)
			continue
		}

		s.msg <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
	}
}

func main() {
	server := NewServer(":3000")
	go func() {
		for msg := range server.msg {
			fmt.Printf("RECV - [%s]: %s\n", msg.from, string(msg.payload))
		}
	}()

	log.Fatal(server.Start())

}
