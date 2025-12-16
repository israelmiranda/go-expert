package main

import (
	"fmt"
)

// Thread 1
func main() {
	forever := make(chan bool)

	// Thread 2
	go func() {
		for i := range 5 {
			fmt.Println(i)
		}
		forever <- true
	}()

	// Thread 1
	<-forever
}
