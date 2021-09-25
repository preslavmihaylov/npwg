package main

import (
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	time.Sleep(5 * time.Second)
}
