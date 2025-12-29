package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func measure(g Geometry) {
	fmt.Printf("Area: %.2f\n", g.Area())
	fmt.Printf("Perimeter: %.2f\n", g.Perimeter())
}

func main() {
	r := &Rectangle{Width: 3, Height: 4}
	c := &Circle{Radius: 5}

	measure(r)
	measure(c)
}
