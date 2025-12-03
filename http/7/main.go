package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jraw := bytes.NewBuffer([]byte(`{"name": "john"}`))
	res, err := c.Post(
		"http://www.google.com",
		"application/json",
		jraw,
	)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
