package main

import (
	"fmt"
	"net"
)

var clientAddr = "127.0.0.1:12344"
var serverAddr = "127.0.0.1:12345"

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

	sendUDP(clientConn, conn.RemoteAddr(), "ping")
	recvUDP(clientConn)
	recvUDP(clientConn)
}

func sendUDP(conn net.PacketConn, addr net.Addr, msg string) {
	_, err := conn.WriteTo([]byte(msg), addr)
	if err != nil {
		panic(err)
	}
}

func recvUDP(conn net.PacketConn) {
	buf := make([]byte, 1024)
	_, recvAddr, err := conn.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("recv from %s: %s\n", recvAddr.String(), string(buf))

}
