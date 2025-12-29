package main

import (
	"github.com/israelmiranda/go-expert/interfaces/3/consumer"
	"github.com/israelmiranda/go-expert/interfaces/3/person"
)

func main() {
	p := person.Person{Name: "John", Age: 32}
	consumer.ProcessPerson(p)
}
