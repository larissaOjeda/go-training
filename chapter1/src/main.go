package main

import "fmt"

func main() {
	a := 1
	b := 2
	sum := add(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)
}

func add(a, b int) int {
	return a + b
}
