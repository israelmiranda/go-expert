package main

import "fmt"

func main() {
	// Memory -> Address -> Value

	// var -> pointer address -> value

	// 1. Declare and initialize an integer variable 'a'
	a := 10

	// 2. Declare a pointer 'p' and assign it the memory address of 'a'
	p := &a

	fmt.Printf("Initial value of a: %d\n", a)
	fmt.Printf("Address of a (&a): %p\n", &a)
	fmt.Printf("Value of p (address of a): %p\n", p)
	fmt.Printf("Value at p (*p): %d\n", *p) // Dereference to get 10

	// 3. Change the value at the address p points to
	*p = 20

	fmt.Printf("\nValue of a after *p = 20: %d\n", a) // Output: 20
}
