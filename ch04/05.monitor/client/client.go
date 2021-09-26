package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Monitor struct {
	*log.Logger
}

func (m *Monitor) Write(p []byte) (n int, err error) {
	return len(p), m.Output(2, string(p))
}

func writerWithPrefix(next io.Writer, prefix string) io.Writer {
	return &PrefixWriter{
		prefix: prefix,
		next:   next,
	}
}

type PrefixWriter struct {
	prefix string
	next   io.Writer
}

func (pw *PrefixWriter) Write(p []byte) (int, error) {
	pp := append([]byte(pw.prefix), p...)
	_, err := pw.next.Write(pp)
	return len(p), err
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1<<19) // 512 KB
	monitor := &Monitor{Logger: log.New(os.Stdout, "", 0)}
	r := io.TeeReader(conn, writerWithPrefix(monitor, "IN: "))
	w := io.MultiWriter(conn, writerWithPrefix(monitor, "OUT: "))
	for {
		_, err = w.Write([]byte("ping"))
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		_, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		fmt.Printf("read [%s]\n", string(buf))
	}
}
