package main

import (
	"example/lib"
	"io"
	"math"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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

	var data lib.Payload
	if math.Abs(float64(rand.Intn(100))) >= 50.0 {
		data = lib.AsBinary([]byte("hello world as bytes!"))
	} else {
		data = lib.AsString([]byte("hello world as string!"))
	}

	_, err := data.WriteTo(conn)
	if err != nil && err != io.EOF {
		if err != io.EOF {
			panic(err)
		}
		return
	}
}
