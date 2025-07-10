package main

import "fmt"

func main() {

	array := [5]int{2, 4, 6, 8, 10}

	jobs := make(chan int, len(array))
	results := make(chan int, len(array))

	for w := 0; w < 3; w++ {
		go func() {
			for num := range jobs {
				results <- num * num
			}
		}()
	}

	for _, num := range array {
		jobs <- num
	}
	close(jobs)

	for i := 0; i < len(array); i++ {
		fmt.Printf("%v ", <-results)
	}
}
