package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("handling incoming request...")
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("result"))
	if err != nil {
		log.Println("couldn't write response: %w", err)
	}
}
