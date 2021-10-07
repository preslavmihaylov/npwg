package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	reqBody := new(bytes.Buffer)
	multipartHeader := createMultipartReq(reqBody)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://httpbin.org/post", reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", multipartHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("status code:", resp.StatusCode)
	fmt.Printf("\n%s\n", b)
}

func createMultipartReq(buf *bytes.Buffer) (multipartHeader string) {
	w := multipart.NewWriter(buf)

	// Add form fields
	for k, v := range map[string]string{
		"date":        time.Now().Format(time.RFC3339),
		"description": "Form values with attached files",
	} {
		err := w.WriteField(k, v)
		if err != nil {
			panic(err)
		}
	}

	// Attach files
	for i, file := range []string{
		"./files/hello.txt",
		"./files/goodbye.txt",
	} {
		filePart, err := w.CreateFormFile(fmt.Sprintf("file%d", i+1),
			filepath.Base(file))
		if err != nil {
			panic(err)
		}

		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(filePart, f)
		_ = f.Close()
		if err != nil {
			panic(err)
		}
	}

	// Finalize multipart request body
	err := w.Close()
	if err != nil {
		panic(err)
	}

	return w.FormDataContentType()
}
