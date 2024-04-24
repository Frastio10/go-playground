package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	addr  string
	ln    net.Listener
	close chan struct{}
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		close: make(chan struct{}),
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

		msg := buf[:n]
		fmt.Println(string(msg))
	}
}

func main() {

	server := NewServer(":3000")
	log.Fatal(server.Start())

}
