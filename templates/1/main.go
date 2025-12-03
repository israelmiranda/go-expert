package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name  string
	Hours int
}

func main() {
	c := Course{Name: "Go", Hours: 40}
	// tmp := template.New("CourseTemplate")
	// tmp, _ = tmp.Parse("Course {{.Name}} - Hours {{.Hours}}")
	// err := tmp.Execute(os.Stdout, c)
	t := template.Must(
		template.New("CourseTemplate").
			Parse("Course {{.Name}} - Hours {{.Hours}}"),
	)
	err := t.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}
