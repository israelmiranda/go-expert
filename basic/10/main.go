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
	sum := func(a, b int) int {
		return sum(a, b)
	}(2, 2)

	multiply := func(n int) int {
		return sum * n
	}

	func() {
		fmt.Println("anonymous")
	}()

	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("multiply: %d\n", multiply(2))
}
