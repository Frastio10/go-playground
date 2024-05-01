package main

func main() {
	server := newServer(":3000")
	server.Start()
}
