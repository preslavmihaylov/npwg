package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var tmpl = template.Must(template.New("hello").Parse("Hello, {{.}}!"))

func main() {
	s := &server{
		log: log.New(os.Stdout, "[server] ", log.Lshortfile),
	}

	srv := &http.Server{
		Addr:              "127.0.0.1:18081",
		Handler:           http.TimeoutHandler(s.DefaultHandler(), 2*time.Minute, ""),
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

type server struct {
	log *log.Logger
}

func (s *server) DefaultHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func(r io.ReadCloser) {
				_, _ = io.Copy(ioutil.Discard, r)
				_ = r.Close()
			}(r.Body)

			var result []byte

			switch r.Method {
			case http.MethodGet:
				s.log.Println("GET invoked")
				result = []byte("friend")
			case http.MethodPost:
				s.log.Println("POST invoked")
				var err error
				result, err = ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Internal server error",
						http.StatusInternalServerError)
					return
				}
			default:
				// not RFC-compliant due to lack of "Allow" header
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			_ = tmpl.Execute(w, string(result))
		},
	)
}
