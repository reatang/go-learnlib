package main

import (
	"fmt"
	"unsafe"
)

// go 1.17 新增了 指针加减操作
// $GOROOT/src/unsafe.go
// func Add(ptr Pointer, len IntegerType) Pointe

// github.com/bigwhite/experiments/tree/master/go1.17-examples/lang/unsafe/add/main.go
const intLen = unsafe.Sizeof(int(0))

func ptrAddOld() {
	var a = [5]int{11, 12, 13, 14, 15}
	for i := 0; i < 5; i++ {
		p := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + uintptr(uintptr(i)*intLen)))
		*p = *p + 10
	}
	fmt.Println(a) // [21 22 23 24 25]
}

// github.com/bigwhite/experiments/tree/master/go1.17-examples/lang/unsafe/add/main.go

func ptrAdd() {
	var a = [5]int{11, 12, 13, 14, 15}
	for i := 0; i < 5; i++ {
		p := (*int)(unsafe.Add(unsafe.Pointer(&a[0]), uintptr(i)*intLen))
		*p = *p + 10
	}
	fmt.Println(a)
}

func main() {
	ptrAddOld()
	ptrAdd()
}
