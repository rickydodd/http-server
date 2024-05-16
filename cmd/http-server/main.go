package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rickydodd/http-server/internal/config"
)

func main() {
	conf, err := config.Config().Build()
	if err != nil {
		log.Fatalln("Error creating config:", err.Error())
	}

	ln, err := net.Listen("tcp",
		// Format `{addr}:{port}`, where {addr} is the address to serve over and {port} is the port to bind to.
		fmt.Sprintf("%s:%d", conf.Addr, conf.Port))
	if err != nil {
		log.Fatalln("Error creating TCP server:", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Error establishing connection:", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\n")
	conn.Close()
}
