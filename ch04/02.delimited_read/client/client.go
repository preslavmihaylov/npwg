package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Printf("recv: %s\n", scanner.Text())
	}

	if scanner.Err() != nil && scanner.Err() != io.EOF {
		panic(err)
	}
}
