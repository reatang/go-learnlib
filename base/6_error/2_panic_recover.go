package main

import (
	"fmt"
	"runtime/debug"
	"sync"
)

func ToPanic() {
	panic("错误")
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Printf("panic recover!\nerr: %s\n", v)
				fmt.Printf("stace:%s", string(debug.Stack()))
			}
		}()

		ToPanic()
		wg.Done()
	}()

	wg.Wait()
}
