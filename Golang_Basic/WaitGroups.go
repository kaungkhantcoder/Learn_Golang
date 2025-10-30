package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		go func(workerID int){
			defer wg.Done()
			worker(workerID)
		}{i}
	}
	wg.Wait()
}
