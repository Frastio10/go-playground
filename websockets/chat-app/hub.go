package main

import "log"

type Hub struct {
	peers map[*Peer]bool

	broadcastCh chan Message

	registerCh chan *Peer

	unregisterCh chan *Peer
}

func (h *Hub) Run() {
	log.Println("Hub is running.")
	for {
		select {
		case peer := <-h.registerCh:
			h.peers[peer] = true

		case peer := <-h.unregisterCh:
			delete(h.peers, peer)
			close(peer.sendCh)

		case message := <-h.broadcastCh:
			for peer := range h.peers {
				if peer != message.from {
					peer.sendCh <- message
				}
			}
		}

	}
}

func newHub() *Hub {
	return &Hub{
		peers:        make(map[*Peer]bool),
		broadcastCh:  make(chan Message),
		registerCh:   make(chan *Peer),
		unregisterCh: make(chan *Peer),
	}
}

func (h *Hub) Unregister(p *Peer) {
	h.unregisterCh <- p
}

func (h *Hub) Register(p *Peer) {
	h.registerCh <- p
}

func (h *Hub) Broadcast(msg Message) {
	h.broadcastCh <- msg
}
