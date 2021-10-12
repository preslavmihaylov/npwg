package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:              "127.0.0.1:18081",
		Handler:           http.TimeoutHandler(DefaultMethodsHandler(), 2*time.Minute, ""),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		panic(err)
	}

	if err = srv.Serve(l); err != http.ErrServerClosed {
		panic(err)
	}

	fmt.Println("server shutdown gracefully...")
}
