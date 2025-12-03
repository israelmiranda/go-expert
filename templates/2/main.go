package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	t := template.Must(
		template.New("template.html").
			ParseFiles("template.html"),
	)
	err := t.Execute(os.Stdout, Courses{
		{Name: "Go", Hours: 40},
		{Name: ".Net", Hours: 30},
		{Name: "Java", Hours: 20},
	})
	if err != nil {
		panic(err)
	}
}
