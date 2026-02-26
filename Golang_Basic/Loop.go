// In Golang, there is only one loop keyword
// Go does not have while, do-while


package main;

import "fmt"

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Go Version of while
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}

// Infinite Loop
for {
	fmt.Println("Running forever")
}

// Using break
for i := 0; i < 10; i++ {
	if i == 5 {
		break
	}
	fmt.Println(i)
}
}

/* for initialization; condition; post {
   code
}
*/
/*
i := 0 => start
i < 5 => condition
i++ => increase

/* output
0
1
2
3
4
*/






