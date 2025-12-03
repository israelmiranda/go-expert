package main

import "fmt"

type ID int

var id ID

func main() {
	fmt.Printf("The type of ID is %T\n", id)
	fmt.Printf("The type of A is %T\n", 1)
	fmt.Printf("The type of B is %T\n", 1.2)

	x := fmt.Sprintf("x is %s", "x")
	fmt.Println(x)
}
