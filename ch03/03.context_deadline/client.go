package main

import (
	"context"
	"net"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := DialCtxWithSleep(ctx, "tcp", "10.0.0.1:http", 5*time.Second)
	if err != nil {
		if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
			panic("couldn't dial server due to timeout")
		} else {
			panic(err)
		}
	}
	defer conn.Close()
}

// Deal.
func DialCtxWithSleep(ctx context.Context, network, address string, sleep time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			time.Sleep(sleep)

			return nil
		},
	}

	return d.DialContext(ctx, network, address)
}
