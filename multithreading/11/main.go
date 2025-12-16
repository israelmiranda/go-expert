package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data <-chan int) {
	for i := range data {
		fmt.Printf("worker %d, received %d\n", workerId, i)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	numberOfWorkers := 2

	// go worker(0, data)
	// go worker(1, data)
	// ...

	for i := range numberOfWorkers {
		go worker(i, data)
	}

	for i := range 10 {
		data <- i
	}
}
