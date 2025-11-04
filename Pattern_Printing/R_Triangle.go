package main

import "fmt"

func R_Triangle() {
	const numRows = 5

	fmt.Println("--- Go Pattern Printer: Right-Angled Triangle ---")

	for i := 1; i <= numRows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}	
}