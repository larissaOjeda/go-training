package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	val, ok := <-ch
	fmt.Printf("val: %d ok: %t\n", val, ok)
}
