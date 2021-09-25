package main

import (
	"context"
	"fmt"
	"net"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	resultCh := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		fmt.Println("spun up fetch goroutine...")
		wg.Add(1)
		go fetch(ctx, &wg, resultCh)
	}

	result := <-resultCh
	cancel()
	wg.Wait()

	close(resultCh)
	fmt.Println(result)
}

func fetch(ctx context.Context, wg *sync.WaitGroup, resultCh chan string) {
	defer wg.Done()

	d := &net.Dialer{}
	conn, err := d.DialContext(ctx, "tcp", "127.0.0.1:12345")
	if err != nil {
		if nerr, ok := err.(net.Error); ok {
			panic(fmt.Sprintf("net error. Error: %s, isTimeout: %v, isTemporary: %v", nerr.Error(), nerr.Timeout(), nerr.Temporary()))
		} else {
			panic(err)
		}
	}
	defer conn.Close()

	res := make([]byte, 1024)
	_, err = conn.Read(res)
	if err != nil {
		panic(err)
	}

	select {
	case <-ctx.Done():
	case resultCh <- string(res):
	}
}
