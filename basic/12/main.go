package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address // embedded
	// Address Address // regular field
}

func main() {
	john := Client{
		Name:   "john",
		Age:    20,
		Active: true,
		Address: Address{
			Street: "Street A",
			Number: 20,
			City:   "City A",
			State:  "AA",
		},
	}
	fmt.Println(john)
}
