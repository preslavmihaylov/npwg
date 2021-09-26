package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var proxyAddr string

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <proxy_address>\n", os.Args[0])
		return
	}

	proxyAddr = os.Args[1]

	fmt.Println("server listening on 127.0.0.1:12345...")
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(source net.Conn) {
	defer source.Close()
	fmt.Println("handling incoming request...")
	defer fmt.Println("closing connection...")

	dest, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	err = proxy(source, dest)
	if err != nil {
		panic(err)
	}
}

func proxy(from io.Reader, to io.Writer) error {
	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	errCh := make(chan error, 2)
	if toIsReader && fromIsWriter {
		go func(errCh chan error) {
			_, err := io.Copy(fromWriter, toReader)
			errCh <- err
		}(errCh)
	}

	go func(errCh chan error) {
		_, err := io.Copy(to, from)
		errCh <- err
	}(errCh)

	return <-errCh
}
