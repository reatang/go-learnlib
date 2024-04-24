package main

import (
	"fmt"
	"unsafe"
)

/*
	slice 转数组
*/

// slice2arrayWithHack 代码中，我们实际上得到是切片底层数组的一份拷贝，修改该拷贝中的元素值，切片中的元素将不会受到影响。
// github.com/bigwhite/experiments/tree/master/go1.17-examples/lang/slice2arrayptr/main.go
func slice2arrayWithHack() {

	// slice
	var b = []int{11, 12, 13}

	// a为b底层数组的拷贝
	var a = *(*[3]int)(unsafe.Pointer(&b[0]))

	// 修改数组内容
	a[1] += 10

	fmt.Printf("%T, %#v\n", a, a)
	fmt.Printf("%T, %#v\n", b, b)
}

// github.com/bigwhite/experiments/tree/master/go1.17-examples/lang/slice2arrayptr/main.go
func slice2arrayptrWithHack() {
	var b = []int{11, 12, 13}

	// p 为 b底层数组的指针
	var p = (*[3]int)(unsafe.Pointer(&b[0]))

	p[1] += 10

	fmt.Printf("%T, %#v\n", p, p)
	fmt.Printf("%T, %#v\n", b, b)
}

// go 1.17 中，语法层面实现了转换
func slice2arrayptr() {
	var b = []int{11, 12, 13}

	// 直接类型转换
	var p = (*[3]int)(b)

	p[1] = p[1] + 10

	fmt.Printf("%T, %#v\n", p, p)
	fmt.Printf("%T, %#v\n", b, b)
}

func main() {
	slice2arrayWithHack()
	fmt.Println("---")
	slice2arrayptrWithHack()

	fmt.Println("===")
	slice2arrayptr()
}
