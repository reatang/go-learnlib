package main

import "fmt"

// 两个一模一样的结构体是否可以自动转换
type AAAA struct {
	A string
	B int
}

type BBBB struct {
	A string
	B int
}

// 结论：可以
func main() {
	a := AAAA{A: "a", B: 1}
	var b BBBB = BBBB(a)

	fmt.Println(b)
}
