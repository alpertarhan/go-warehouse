package main

import (
	"log"
	"net"
)

func main() {
	s := newServer() // initial server

	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server : %s", err.Error())
	}

	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)
	log.Printf("start server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}

}
