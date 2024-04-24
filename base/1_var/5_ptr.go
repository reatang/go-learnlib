package main

import "fmt"

// 活用指针

type PA struct {
	Name string
}

func main() {

	// 1、指针
	aptr := &PA{"rea"}
	bptr := aptr
	bptr.Name = "reatang"
	bptr = nil
	fmt.Println(aptr.Name)

	// 2、拆解写法，两个指针指向同一个内存地址。
	var aptr2 *PA
	aptr2 = &PA{"rea"}
	var bptr2 *PA
	bptr2 = aptr2
	bptr2.Name = "reatang"
	fmt.Println(aptr2.Name)

	// 3、指针传递
	ptr1()
}

type P1 struct {
	Name string
}

// 测试指针
func ptr1() {
	p := &P1{"ptr1"}
	// 打印指针地址和指向的地址
	fmt.Printf("p1: %p, %p\n", &p, p)
	ptr2(p)
	fmt.Printf("p1: %p, %p\n", &p, p)
	fmt.Println(p.Name)
}

// 修改指针指向的内存中的内容
func ptr2(p *P1) {
	p2 := &P1{"ptr2"}
	fmt.Printf("p2: %p, %p\n", &p2, p2)
	*p = *p2
}
