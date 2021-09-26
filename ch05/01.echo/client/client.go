package main

import (
	"fmt"
	"net"
)

var serverAddr = "127.0.0.1:12345"

func main() {
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := []byte("ping")
	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("recv from %s: %s\n", conn.RemoteAddr().String(), string(buf))
}
