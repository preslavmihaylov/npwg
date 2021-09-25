package main

import (
	"fmt"
	"log"
	"net"
	"time"
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

	err := conn.SetDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		panic(err)
	}

	result := make([]byte, 1024)
	_, err = conn.Read(result)
	if err != nil {
		log.Printf("couldn't read response: %v", err)
		return
	}

	fmt.Println("received input: " + string(result))
}
