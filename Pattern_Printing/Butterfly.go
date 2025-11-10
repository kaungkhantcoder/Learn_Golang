package main

import (
	"fmt"
)

func printButterfly(n int) {

	for i := 1; i <= n; i++ {

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		gapSpaces := 2 * (n - i)
		for j := 0; j < gapSpaces; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		gapSpaces := 2 * (n - i)
		for j := 0; j < gapSpaces; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func Butterfly() {
	const patternSize = 5

	fmt.Printf("Generating Butterfly Pattern of size %d:\n\n", patternSize)
	printButterfly(patternSize)
}