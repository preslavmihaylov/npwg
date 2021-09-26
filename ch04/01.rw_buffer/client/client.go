package main

import (
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

	buf := make([]byte, 1<<19) // 512 KB
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		fmt.Printf("read %d bytes\n", n)
	}
}
