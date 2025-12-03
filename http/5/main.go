package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Blog"))
	})

	fmt.Println("starting server on :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal("server failed: ", err)
	}
}
