package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := range 5 {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	task("A")
	task("B")
}
