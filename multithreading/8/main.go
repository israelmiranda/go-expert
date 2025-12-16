package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)

	// Thread 2
	go publish(ch)

	// Thread 1
	reader(ch)
}

func reader(ch <-chan int) {
	for i := range ch {
		fmt.Printf("received: %d\n", i)
	}
}

func publish(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}
