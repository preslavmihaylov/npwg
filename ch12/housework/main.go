package main

import (
	"flag"
	"fmt"
	"housework/controller"
	"housework/server"
	"housework/storage"
	"housework/storage/gob"
	"housework/storage/json"
	"housework/storage/protobuf"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dataFile string
var format string

func init() {
	flag.StringVar(&dataFile, "file", "housework.db", "data file")
	flag.StringVar(&format, "format", "json", "db storage format. Options: json, gob, protobuf")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`Usage: %s [flags] [add chore, ...|complete #]
    add         add comma-separated chores
    complete    complete designated chore
	serve       serve grpc housework server
Flags:
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if strings.ToLower(flag.Arg(0)) == "serve" {
		err := server.Serve("localhost:34443")
		if err != nil {
			log.Fatalf("received error from server: %v", err)
		}

		return
	}

	var storage storage.Storage
	switch format {
	case "json":
		storage = json.NewStorage()
	case "gob":
		storage = gob.NewStorage()
	case "protobuf":
		storage = protobuf.NewStorage()
	default:
		panic("unknown storage format")
	}

	houseworkCtrl := controller.New(storage, dataFile)

	var err error
	switch strings.ToLower(flag.Arg(0)) {
	case "add":
		err = houseworkCtrl.Add(strings.Join(flag.Args()[1:], " "))
	case "complete":
		err = houseworkCtrl.Complete(flag.Arg(1))
	}

	if err != nil {
		log.Fatal(err)
	}

	err = houseworkCtrl.List()
	if err != nil {
		log.Fatal(err)
	}
}
