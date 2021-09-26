package main

import (
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("hello there mate bate faith"))
	if err != nil && err != io.EOF {
		if err != io.EOF {
			panic(err)
		}
		return
	}
}
