package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	resp, err := http.Head("https://www.time.gov/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("Date")
	if date == "" {
		panic("no Date header received from time.gov")
	}

	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		panic(err)
	}

	fmt.Printf("time.gov: %s (skew %s)\n", dt, now.Sub(dt))
}
