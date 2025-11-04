package main

import "fmt"

func Pyramid() {
	const numRows = 5

	fmt.Println("--- Go Pattern Printer: Full Pyramid ---")

	for i := 1; i <= numRows; i++ {
		for j := 1; j <= numRows - i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}