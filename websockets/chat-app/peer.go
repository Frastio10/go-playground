package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Peer struct {
	hub    *Hub
	conn   *websocket.Conn
	sendCh chan Message
}

func (p *Peer) ReadLoop() {
	defer func() {
		p.hub.Unregister(p)
		p.conn.Close()
	}()

	for {
		_, buf, err := p.conn.ReadMessage()
		if err != nil {
			log.Println("Read Error", err)
			break
		}

		log.Printf("Recieved - [%s]: %s\n", p.conn.RemoteAddr().String(), string(buf))

		p.hub.Broadcast(Message{from: p, payload: buf})
	}

}

func (p *Peer) WriteLoop() {
	defer func() {
		p.conn.Close()
	}()

	for {
		select {
		case msg := <-p.sendCh:
			p.conn.WriteMessage(websocket.TextMessage, msg.payload)
		}
	}

}

func (p *Peer) HandleConnection() {
	p.hub.Register(p)

	log.Println("New client connected: ", p.conn.RemoteAddr().String())

	go p.ReadLoop()
	go p.WriteLoop()
}
