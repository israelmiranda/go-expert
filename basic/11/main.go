package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	john := Client{
		Name:   "john",
		Age:    20,
		Active: true,
	}
	john.Active = false
	fmt.Println(john)
}
