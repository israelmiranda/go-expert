package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("Hello, world!")
	size, err := f.Write([]byte("Hello, world!"))
	if err != nil {
		panic(err)
	}
	log.Printf("File created with size: %d bytes\n", size)
	f.Close()

	// read
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// read buffer
	c, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(c)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	c.Close()
}
