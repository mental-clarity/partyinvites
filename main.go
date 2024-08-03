package main

import (
	"fmt"
	"text/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var resonses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTempletes() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err != nil {
			panic(err)
		}
		templates[name] = t
		fmt.Println("Loaded template", index, name)
	}
}

func main() {
	loadTempletes()
}
