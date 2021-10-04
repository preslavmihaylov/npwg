package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
)

var socketDir = "/tmp/echo_unix"

func main() {
	socket, teardown, err := createUnixSocket()
	if err != nil {
		panic(err)
	}
	defer teardown()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		cancel()
	}()

	prestart := func() error {
		return os.Chmod(socket, os.ModeSocket|0666)
	}

	if err := runServer(ctx, "unix", socket, prestart); err != nil {
		panic(err)
	}
}

func createUnixSocket() (string, func(), error) {
	err := os.Mkdir(socketDir, 0766)
	if err != nil {
		return "", nil, fmt.Errorf("couldn't mkdir: %w", err)
	}
	teardown := func() {
		fmt.Println("tearing down...")
		os.RemoveAll(socketDir)
	}

	socket := filepath.Join(socketDir, "main.sock")
	return socket, teardown, nil
}

func runServer(ctx context.Context, network string,
	addr string, prestart func() error) error {
	s, err := net.Listen(network, addr)
	if err != nil {
		return fmt.Errorf("binding to %s %s: %w", network, addr, err)
	}

	err = prestart()
	if err != nil {
		return fmt.Errorf("received error from server pre-start hook: %w", err)
	}

	go func() {
		<-ctx.Done()
		_ = s.Close()
	}()

	for {
		conn, err := s.Accept()
		if err != nil {
			return err
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		fmt.Printf("recv: [%s]\n", string(buf))
		_, err = conn.Write(buf[:n])
		if err != nil {
			return
		}
	}
}
