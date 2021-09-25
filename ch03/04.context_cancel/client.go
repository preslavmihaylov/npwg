package main

import (
	"context"
	"fmt"
	"net"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	conn, err := DialCtxWithSleep(ctx, "tcp", "10.0.0.1:http", 5*time.Second)
	if err != nil {
		if nerr, ok := err.(net.Error); ok {
			panic(fmt.Sprintf("net error. Error: %s, isTimeout: %v, isTemporary: %v", nerr.Error(), nerr.Timeout(), nerr.Temporary()))
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
