package main

import "fmt"

func receive(greeting string, ch chan<- string) {
	ch <- greeting
}

func read(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go receive("hello", ch)
	read(ch)
}
