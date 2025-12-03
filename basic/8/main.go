package main

import (
	"errors"
	"fmt"
	"log"
)

func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("undefined")
	}
	return a / b, nil
}

func show(total int) {
	fmt.Printf("total: %d\n", total)
}

func main() {
	show(sum(2, 2))

	r, err := divide(2, 0)
	if err != nil {
		log.Printf("invalid param: %s", err)
	}
	show(r)
}
