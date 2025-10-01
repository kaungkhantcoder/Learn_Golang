package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID 		int
	ID			string	`json:"id"`
	Title		string	`json:"title"`
	Completed	bool	`json:"completed"`
}

func main() {
	
	url := "http://localhost:3000/todos/1"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		todoItem := Todo{}

		decoder := json.NewDecoder(response.Body)

		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal("Decode error: ", err)
		}

		fmt.Printf("Data from API: %+v", todoItem)

	}
}