package main

import (
	"net"
	"syscall"
	"time"
)

func main() {
	conn, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err != nil {
		if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
			panic("couldn't dial server due to timeout")
		} else {
			panic(err)
		}
	}
	defer conn.Close()
}

func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        address,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}

	return d.Dial(network, address)
}
