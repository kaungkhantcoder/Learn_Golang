package main

import (
	"fmt" // Format Module
	"html/template"
	
	"log"
	"net/http"
)

type Club struct {
	Name 	string
	City 	string
}

func main() {
	fmt.Println("Hello World")

	h1 := func (w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		clubs := map[string][]Club{
			"Club": {
				{Name: "Arsenal", City: "London"},
				{Name: "Chelsea", City: "London"},
				{Name: "Manchester United", City: "Manchester"},
			},
		}
		tmpl.Execute(w, clubs)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}