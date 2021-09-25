package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const defaultPingInterval = 2 * time.Second

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("handling incoming request...")
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resetCh := make(chan time.Duration, 1)
	go pinger(ctx, conn, resetCh)
	for {
		result := make([]byte, 1024)
		_, err := conn.Read(result)
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Println("closing connection...")
			default:
				log.Printf("couldn't read response: %v", err)
			}

			return
		}

		fmt.Println("received input: " + string(result))
		resetCh <- 0
	}
}

func pinger(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	interval := defaultPingInterval
	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newInterval := <-reset:
			if !timer.Stop() {
				<-timer.C
			}

			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C:
			if _, err := w.Write([]byte("ping")); err != nil {
				// track and act on consecutive timeouts here
				return
			}
		}

		_ = timer.Reset(interval)
	}
}
