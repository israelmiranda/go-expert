package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Printf("sum: %d\n", sum(5, 4, 3, 2, 1))
}
