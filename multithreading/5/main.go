package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan string) // empty channel

	// Thread 2
	go func() {
		ch <- "Hello Channel!" // channel full
	}()

	// Thread 1
	msg := <-ch // channel is empty
	fmt.Println(msg)
}
