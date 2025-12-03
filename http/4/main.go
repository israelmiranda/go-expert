package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	channel := make(chan struct{})

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World 1!"))
	})

	go func() {
		fmt.Println("Starting Server 1 on :3000...")
		if err := http.ListenAndServe(":3000", mux1); err != nil {
			log.Fatal("Server 1 failed: ", err)
		}
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World 2!"))
	})

	go func() {
		fmt.Println("Starting Server 2 on :3001...")
		if err := http.ListenAndServe(":3001", mux2); err != nil {
			log.Fatal("Server 2 failed: ", err)
		}
	}()

	<-channel
}
