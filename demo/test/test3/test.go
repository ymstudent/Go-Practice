package main

import (
	"fmt"
	"sync"
)

func main() {
	s1 := make([]int, 0)
	for i := 0; i < 100; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(s1)
	s2 := make([]int, len(s1))
	wg := sync.WaitGroup{}
	for k, v := range s1 {
		wg.Add(1)
		go func() {
			s2[k] = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(s2)
}
