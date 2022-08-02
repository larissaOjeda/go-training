package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d\n", i)
			time.Sleep(1 * time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Print("Finish\n")
}
