package main

import (
	"fmt"

	"github.com/isralmiranda/go-expert/packages/2/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Sum())
}
