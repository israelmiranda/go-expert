package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(5)

	// Thread 2
	go publish(ch)

	// Thread 3
	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch <-chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("received: %d\n", i)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}
