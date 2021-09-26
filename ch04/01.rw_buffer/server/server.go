package main

import (
	"crypto/rand"
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

	payload := make([]byte, 1<<24) // 16 MB
	_, err := rand.Read(payload)   // generate a random payload
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
		return
	}

	_, err = conn.Write(payload)
	if err != nil && err != io.EOF {
		if err != io.EOF {
			panic(err)
		}
		return
	}
}
