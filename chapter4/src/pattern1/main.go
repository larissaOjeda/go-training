package main

import (
	"fmt"
)

func gen(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func sq(ch <-chan int, chRes chan<- int) {
	for v := range ch {
		chRes <- v * v
	}
	close(chRes)
}

func main() {
	numbers := make(chan int, 10)
	results := make(chan int, 10)
	go gen(numbers)
	go sq(numbers, results)

	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Print("Finished\n")
}
