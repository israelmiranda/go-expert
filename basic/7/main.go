package main

import "fmt"

func main() {
	incomes := map[string]int{
		"john": 1000,
		"jane": 2000,
		"jake": 3000,
	}

	fmt.Println(incomes["john"])
	delete(incomes, "john")
	incomes["josh"] = 4000

	value, ok := incomes["israel"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found")
	}

	fmt.Println(incomes)

	incomes1 := make(map[string]int)
	incomes2 := map[string]int{}

	fmt.Println(incomes1)
	fmt.Println(incomes2)

	incomes3 := map[string]int{
		"john": 5000,
		"jake": 6000,
		"jane": 7000,
	}

	for k, v := range incomes3 {
		fmt.Printf("key: %s - value: %d\n", k, v)
	}

	for _, v := range incomes3 { // blanck identifier
		fmt.Printf("value: %d\n", v)
	}
}
