package main

func main() {
	var value any = "John Doe"
	println(value.(string))

	res, ok := value.(int)
	if ok {
		println(res)
	}
}
