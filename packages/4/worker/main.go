package main

import (
	"fmt"

	"github.com/israelmiranda/go-expert/packages/4/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Sum())
}
