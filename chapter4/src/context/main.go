package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chDone := make(chan struct{})
	ch := make(chan int, 1000)
	ctx, cnl := context.WithCancel(context.Background())

	go func() {
		i := 0
		for {
			if ctx.Err() != nil {
				fmt.Print("Producer Done\n")
				chDone <- struct{}{}
				return
			}
			ch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Print("Writer Done\n")
				chDone <- struct{}{}
				return
			case n := <-ch:
				fmt.Printf("%d\n", n)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	cnl()
	<-chDone
	<-chDone
	fmt.Print("Finish\n")
}
