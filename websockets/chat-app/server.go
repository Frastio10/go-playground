package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

var upgrader = websocket.Upgrader{}

func (s *Server) ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	http.ServeFile(w, r, "index.html")
}

func (s *Server) ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Serve WS: ", err)
	}

	peer := &Peer{hub: hub, conn: conn, sendCh: make(chan Message)}
	peer.HandleConnection()
}

func (s *Server) Start() {
	hub := newHub()
	go hub.Run()

	http.HandleFunc("/", s.ServeHome)
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		s.ServeWs(hub, w, r)
	})

	server := &http.Server{
		Addr:              s.addr,
		ReadHeaderTimeout: 3 * time.Second,
	}

	s.httpServer = server

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Listen and Serve: ", err)

	}
}

func newServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}
