package main

import (
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("hello world 1"))
	conn.Write([]byte("hello world 2"))
	conn.Write([]byte("hello world 3"))
}
