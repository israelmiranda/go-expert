package main

import "fmt"

// by value copy
// func sum(a, b int) int {
// 	a = 10
// 	return a + b
// }

// func main() {
// 	n1 := 2
// 	n2 := 3

// 	fmt.Printf("sum: %d\n", sum(n1, n2))

// 	fmt.Printf("value of n1: %d\n", n1)
// }

// by ref
// func sum(a, b *int) int {
// 	*a = 10
// 	return *a + *b
// }

// func main() {
// 	n1 := 2
// 	n2 := 3

// 	fmt.Printf("sum: %d\n", sum(&n1, &n2))

// 	fmt.Printf("value of n1: %d\n", n1)
// }

func modString(a *string) {
	*a = "modified"
	fmt.Println("Inside: ", *a)
}

func main() {
	a := "hello"

	fmt.Println("Before: ", a)
	modString(&a)
	fmt.Println("After: ", a)
}
