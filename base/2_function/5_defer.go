package main

import "fmt"

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	
	return 1
}

func f2() (result int) {
	defer func() {
		result++
	}()
	return 0
}


func main() {
	r := f()
	fmt.Println(r)
	
	r2 := f2()
	fmt.Println(r2)
}
