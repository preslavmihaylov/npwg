package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	First, Last string
}

func main() {
	buf := new(bytes.Buffer)
	u := User{First: "Adam", Last: "Woodbeck"}
	err := json.NewEncoder(buf).Encode(&u)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080", "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status code:", resp.StatusCode)
}
