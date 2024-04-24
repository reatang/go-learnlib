package main

import (
	"fmt"
	"sync"
)

/*

silce 并发操作

*/

func main() {
	s1 := make([]int, 1000)

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i, _ := range s1 {
		go func(i int) {
			s1[i] = i + 100

			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < 1000; i++ {
		fmt.Println(s1[i])
	}
}
