package main

import (
	"fmt"
	"net"
)

var clientAddr = "127.0.0.1:12346"
var serverAddr = "127.0.0.1:12344"

func main() {
	clientConn, err := net.ListenPacket("udp", clientAddr)
	if err != nil {
		panic(fmt.Sprintf("binding to udp %s: %v", clientAddr, err))
	}
	defer clientConn.Close()

	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	sendUDP(clientConn, conn.RemoteAddr(), "interrupt")
}

func sendUDP(conn net.PacketConn, addr net.Addr, msg string) {
	_, err := conn.WriteTo([]byte(msg), addr)
	if err != nil {
		panic(err)
	}
}
