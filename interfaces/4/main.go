package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func describe(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Printf("String value: %s\n", str)
		return
	}

	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	case Rectangle:
		fmt.Printf("Rectangle with area: %.2f\n", v.Area())
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	describe("hello")
	describe(42)
	describe(true)
	describe(Rectangle{5, 7})
	describe(3.14)
}
