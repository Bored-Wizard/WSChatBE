package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()
	go s.clean()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("unable to start server :8888")
	}
	defer listener.Close()
	log.Printf(("started server on :8888"))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
