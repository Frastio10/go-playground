package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Couldn't start a listener: ", err)
	}

	log.Println("Waiting for connection..")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Couldn't start a connection", err)
			continue
		}

		// handleConn(conn) // can only accept one connection at a time
		go handleConn(conn) // handle every connection on a single goroutine

	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	log.Println("Connection established.")

	// runs on every tick
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}

}
