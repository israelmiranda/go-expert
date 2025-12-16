package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type message struct {
	id      int64
	message string
}

func main() {
	c1 := make(chan message)
	c2 := make(chan message)

	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			c1 <- message{i, "Hello from RabbitMQ"}
		}
	}()
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			c2 <- message{i, "Hello from Kafka"}
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Receive message from RabbitMQ: %d - %s\n", msg.id, msg.message)
		case msg := <-c2:
			fmt.Printf("Receive message from Kafka: %d - %s\n", msg.id, msg.message)
		case <-time.After(3 * time.Second):
			println("timeout")
		}
	}
}
