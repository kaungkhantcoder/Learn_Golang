package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Todo struct {
	UserID 		int		`json:"-"`
	ID			int		`json:"-"`
	Title		string	`json:"title,omitempty"`
	Completed	bool	`json:"completed"`
}

func main() {

	todoItem := &Todo{1, 1, "", false}
	
	// Convert back to JSON
	todo, err := json.MarshalIndent(todoItem, "", "\t")
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println(string(todo))
}