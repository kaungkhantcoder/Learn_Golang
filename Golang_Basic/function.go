package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(5, 10)
	fmt.Println(res)

	res = plusPlus(5, 10, 15)
	fmt.Println(res)

	switch_fun()
}