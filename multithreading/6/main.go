package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan string) // empty channel

	// Thread 2
	go func() {
		// channel full
		for i := range 5 {
			ch <- fmt.Sprintf("Hello Channel %d!", i)
		}
	}()

	// Thread 1
	for i := range 5 {
		msg := <-ch // channel is empty
		fmt.Printf("%d: %s\n", i, msg)
	}
}
