package main

type ID int
type Message string

var (
	id      ID      = 1
	message Message = "message"
)

func main() {
	println(id)
	println(message)
}
