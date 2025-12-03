package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(len(arr))
	fmt.Println(arr)

	for i := 0; i <= len(arr)-1; i++ {
		fmt.Println(i)
	}

	for i := range arr {
		fmt.Println(i)
	}

	for i, v := range arr {
		fmt.Printf("i: %d - v: %d\n", i, v)
	}
}
