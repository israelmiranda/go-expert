package main

import (
	"fmt"
	"strings"
)

func filter[T int | string](arr []T, condition func(T) bool) []T {
	var result []T
	for _, element := range arr {
		if condition(element) {
			result = append(result, element)
		}
	}
	return result
}

type CustomNumber int

type Number interface {
	~int | float64
}

// type Number interface {
// 	int | float64
// }

func sum[T Number](values ...T) T {
	var total T = 0
	for _, v := range values {
		total += v
	}
	return total
}

func compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	f1 := filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Printf("filter numbers: %v\n", f1)

	names := []string{"john", "jake", "jane", "lara"}
	f2 := filter(names, func(s string) bool {
		return strings.Contains(s, "l")
	})

	fmt.Printf("filter names: %v\n", f2)

	fmt.Printf("sum int: %d\n", sum(1, 2, 3, 4))
	fmt.Printf("sum float: %v\n", sum(1.1, 2.2, 3.3, 4.4))

	var n1 CustomNumber = 1
	var n2 CustomNumber = 1
	fmt.Printf("sum custom type: %d\n", sum(n1, n2))

	fmt.Printf("compare: %v\n", compare(2, 2))
}
