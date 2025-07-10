package main

import "fmt"

func main() {

	result := make(chan int, 5)
	array := [5]int{2, 4, 6, 8, 10}

	for _, value := range array {
		go func(value int) {
			result <- value * value
		}(value)
	}

	for i := 0; i < len(array); i++ {
		fmt.Printf("%v ", <-result)
	}

}
