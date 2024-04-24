package main

import "fmt"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// 对int类型做二进制操作
func int2bin() {
	mark := 0b1111

	var b int
	for _, v := range makeRange(0, 127) {
		b = v & mark
		fmt.Printf("%-8b %-4b %d\n", v, b, b)
	}
}

func main() {
	int2bin()
}
