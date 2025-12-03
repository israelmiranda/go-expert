package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Client) Deactivate() {
	c.Active = false
	fmt.Println(c)
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

	john.Deactivate()
	fmt.Println(john)
}
