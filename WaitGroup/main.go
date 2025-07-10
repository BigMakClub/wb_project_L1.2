package main

import (
	"fmt"
	"sync"
)

func worker(id int, value int) {
	fmt.Printf("Worker %d starting\n", id)
	fmt.Printf("result %d\n", value*value)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	array := [5]int{2, 4, 6, 8, 10}

	for i, value := range array {
		wg.Add(1)
		go func(i, value int) {
			defer wg.Done()
			worker(i, value)
		}(i, value)
	}

	wg.Wait()
}
