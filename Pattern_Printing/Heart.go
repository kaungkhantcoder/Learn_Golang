package main

import (
	"fmt"
)


func PrintHeartPattern(size int) {

	if size%2 == 0 {
		size++
	}

	for row := size / 2; row <= size; row += 2 {

		for col := 1; col < size-row; col += 2 {
			fmt.Print(" ")
		}


		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}

	
		for col := 1; col <= size-row; col++ {
			fmt.Print(" ")
		}


		for col := 1; col <= row-1; col++ {
			fmt.Print("*")
		}

		fmt.Println() 
	}


	for row := size; row >= 0; row-- {

		for col := row; col < size; col++ {
			fmt.Print(" ")
		}


		for col := 1; col <= (row*2)-1; col++ {
			fmt.Print("*")
		}

		if (row*2)-1 > 0 {
			fmt.Println()
		}
	}
}

func Heart() {

	heartSize := 5
	fmt.Printf("Printing a heart pattern with size %d:\n\n", heartSize)
	PrintHeartPattern(heartSize)
}