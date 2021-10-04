package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
)

var socketType = "unixgram"
var socketDir = "/tmp/echo_unixgram"

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
		return os.Chmod(socket, os.ModeSocket|0622)
	}

	if err := runServer(ctx, socketType, socket, prestart); err != nil {
		panic(err)
	}
}

func createUnixSocket() (string, func(), error) {
	err := os.Mkdir(socketDir, 0766)
	if err != nil {
		return "", nil, fmt.Errorf("couldn't mkdir: %w", err)
	}
	teardown := func() { os.RemoveAll(socketDir) }

	socket := filepath.Join(socketDir, "server.sock")
	return socket, teardown, nil
}

func runServer(ctx context.Context, network string,
	addr string, prestart func() error) error {
	s, err := net.ListenPacket(network, addr)
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
		input := make([]byte, 1024)
		_, dest, err := s.ReadFrom(input)
		if err != nil {
			return err
		}

		go handleRequest(s, dest, input)
	}
}

func handleRequest(conn net.PacketConn, dest net.Addr, input []byte) {
	fmt.Printf("recv: [%s]\n", string(input))
	_, err := conn.WriteTo(input, dest)
	if err != nil {
		log.Printf("received err when sending packet: %v", err)
	}
}
