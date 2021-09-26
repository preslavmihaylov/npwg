package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("server listening on 127.0.0.1:12346...")
	listener, err := net.Listen("tcp", "127.0.0.1:12346")
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
	fmt.Println("handling incoming request...")
	defer fmt.Println("closing connection...")

	for i := 0; i < 4; i++ {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			return
		}

		_, err = conn.Write([]byte("pong"))
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			return
		}

		fmt.Printf("read [%s]\n", string(buf))
	}
}
