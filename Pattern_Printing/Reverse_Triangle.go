package main

import (
	"fmt"
)

func ReverseTriangle() {
	rows := 5

	fmt.Println("Reverse Triangle Pattern:")
	for i := 0; i < rows; i++ {

		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*(rows-i)-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}