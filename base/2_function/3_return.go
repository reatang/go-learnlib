package main

import "fmt"

// 多值返回

// 普通的函数定义
func add(x, y int) int {
	return x + y
}

// 1、多值返回
func reverse(a, b int) (int, int) {
	return b, a
}

// 2、预定义返回变量的名称 多值返回
func reverse2(a, b int) (rb int, ra int) {

	rb = a
	ra = b
	return
}

// 3、预定义返回变量的名称 同样类型可以合并类型名称
func reverse3(a, b int) (rb, ra int) {

	rb = a
	ra = b
	return
}

func main() {
	fmt.Println(add(1, 2))

	// 普通的多值返回
	a, b := reverse(1, 2)
	fmt.Println(a, b)

	a, b = reverse2(3, 4)
	fmt.Println(a, b)

	a, b = reverse3(5, 6)
	fmt.Println(a, b)
}
