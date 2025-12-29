// In the package with the implementation
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type PersonReader interface {
	GetName() string
	GetAge() int
}

func (p Person) GetName() string { return p.Name }
func (p Person) GetAge() int     { return p.Age }

func ProcessPerson(p PersonReader) {
	fmt.Printf("%s - %d\n", p.GetName(), p.GetAge())
}

func main() {
	p := Person{Name: "John", Age: 32}
	ProcessPerson(p)
}
