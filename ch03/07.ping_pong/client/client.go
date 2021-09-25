package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	waitForPings(conn, 4)
	sendPongsWithCadence(conn, 4, 1*time.Second)
	waitForPings(conn, 4)
}

func waitForPings(conn net.Conn, cnt int) {
	buf := make([]byte, 1024)
	for i := 0; i < cnt; i++ {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		log.Printf("recv: %s", buf[:n])
	}
}

func sendPongsWithCadence(conn net.Conn, cnt int, cadence time.Duration) {
	for i := 0; i < cnt; i++ {
		_, err := conn.Write([]byte("pong"))
		if err != nil {
			panic(err)
		}

		time.Sleep(cadence)
	}
}
