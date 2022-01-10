package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:8090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go requestProcess(conn)
	}
}

func requestProcess(conn net.Conn) {
	io.Copy(os.Stdout, conn)
}
