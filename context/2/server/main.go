package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("failed to start on port 3000")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Init request.")
	defer log.Println("Finished.")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request has been processed successfully")
		w.Write([]byte("Request has been processed successfully"))
	case <-ctx.Done():
		log.Println("Request was cancelled by client")
	}
}
