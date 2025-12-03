package main

import "fmt"

// func showType(value interface{}) {
func showType(value any) {
	fmt.Printf("The type of the value is %T, and the value is %v\n", value, value)
}

func main() {
	// var x interface{} = 10
	// var y interface{} = "Hello, world!"
	var x any = 10
	var y any = "Hello, world!"

	showType(x)
	showType(y)
}
