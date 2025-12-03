package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/israelmiranda/go-expert/basic/21/math"
)

func main() {
	s := math.Sum(1, 1)

	fmt.Printf("sum: %v\n", s)
	fmt.Printf("const: %v\n", math.PI)

	fmt.Println(uuid.New())
}
