package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID 		int 	`json:"userId"`
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Completed	bool	`json:"completed"`
}

func main() {
	
	url := "https://dummyjson.com/todos/1"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		todoItem := Todo{}

		json.NewDecoder(response.Body).Decode(&todoItem)

		fmt.Printf("Data from API: %+v", todoItem)

	}
}