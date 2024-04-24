package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 3, 5)

	n := copy(s2, s1[5:])

	fmt.Printf("%d, %+v\n", n, s2)

	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
