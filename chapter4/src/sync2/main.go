package main

import (
	"fmt"
	"sync"
)

type sum struct {
	sync.Mutex
	sum int
}

func (s *sum) add(i int) {
	s.Lock()
	defer s.Unlock()
	s.sum = s.sum + i
}

func main() {
	wg := sync.WaitGroup{}
	s := sum{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.add(i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Sum: %d\n", s.sum)
}
