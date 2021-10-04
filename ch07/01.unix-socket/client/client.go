package main

import (
	"fmt"
	"net"
)

var socketDest = "/tmp/echo_unix/main.sock"

func main() {
	conn, err := net.Dial("unix", socketDest)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < 3; i++ {
		msg := []byte("ping")
		_, err = conn.Write([]byte(msg))
		if err != nil {
			panic(err)
		}
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("recv from %s: [%s]\n", conn.RemoteAddr().String(), string(buf))
}
