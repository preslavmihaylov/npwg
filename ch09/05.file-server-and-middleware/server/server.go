package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			RestrictPrefix(".", http.FileServer(http.Dir("./files"))),
		),
	)

	srv := &http.Server{
		Addr:              "127.0.0.1:18081",
		Handler:           mux,
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
