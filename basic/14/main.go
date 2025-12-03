package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person interface {
	Deactivate()
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

type Company struct {
	Name string
}

func (c Company) Deactivate() {
	fmt.Println("not implement")
}

func Deactivate(person Person) {
	person.Deactivate()
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

	Deactivate(john)
	fmt.Println(john)

	comp := Company{Name: "My Company"}
	Deactivate(comp)
	fmt.Println(comp)
}
