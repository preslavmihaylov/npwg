package main

import (
	"fmt"
	"net"
)

var addr = "127.0.0.1:12345"

func main() {
	packetConn, err := net.ListenPacket("udp", addr)
	if err != nil {
		panic(fmt.Sprintf("error binding to udp %s: %v", addr, err))
	}
	defer packetConn.Close()

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := packetConn.ReadFrom(buf) // client to server
		if err != nil {
			return
		}

		fmt.Printf("recv from %s: %s\n", clientAddr.String(), string(buf))
		_, err = packetConn.WriteTo(buf[:n], clientAddr) // server to client
		if err != nil {
			return
		}
	}
}
