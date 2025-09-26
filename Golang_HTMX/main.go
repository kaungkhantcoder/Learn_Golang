package main

import (
	"fmt" // Format Module
	"html/template"
	"time"

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

	h2 := func (w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		club := r.PostFormValue("club")
		city := r.PostFormValue("city")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "club-name", Club{Name: club, City: city})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-club/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}