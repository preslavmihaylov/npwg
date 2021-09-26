package main

import (
	"flag"
	"io/ioutil"
	"log"
	"tftp-server/server"
	"time"
)

var (
	address = flag.String("a", "127.0.0.1:69", "listen address")
	payload = flag.String("p", "payload.svg", "file to serve to clients")
)

func main() {
	flag.Parse()

	p, err := ioutil.ReadFile(*payload)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.New(p, 5, 6*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.ListenAndServe(*address))
}
