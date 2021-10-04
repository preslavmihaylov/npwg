package main

import (
	"fmt"
	"net"
	"os"
)

var socketSrc = "/tmp/echo_unixgram/client.sock"
var socketDest = "/tmp/echo_unixgram/server.sock"

func main() {
	conn, err := net.ListenPacket("unixgram", socketSrc)
	if err != nil {
		panic(err)
	}
	defer func() {
		os.Remove(socketSrc)
		conn.Close()
	}()

	dest := &net.UnixAddr{
		Name: socketDest,
		Net:  "unixgram",
	}
	for i := 0; i < 3; i++ {
		msg := []byte("ping")
		_, err = conn.WriteTo([]byte(msg), dest)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 3; i++ {
		buf := make([]byte, 1024)
		_, serverAddr, err := conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("recv from %s: [%s]\n", serverAddr.String(), string(buf))
	}
}
