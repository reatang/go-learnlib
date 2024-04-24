package main

import "fmt"

// 1、常量
const SOME_CONST = 666

// 2、批量
const (
	C1 = 0
	C2 = 1
	C3 = 2
)

// 3、使用 iota 从0按行累加1，iota 只能用于 const() 内
const (
	C_A_1 = iota // iota = 0
	C_A_2        // iota = 1
	C_A_3        // iota = 2
	// 空行不算
	C_A_4                             // iota = 3
	C_A_5, C_A_6 = iota + 1, iota + 2 // iota = 4, C_A_5 = 5，C_A_6 = 6
	C_A_7, C_A_8                      // iota = 5, C_A_5 = 6，C_A_6 = 7
)

func main() {

	// 局部常量
	const SOME_CONST2 = "汉字"

	fmt.Println("123", SOME_CONST, SOME_CONST2)

	fmt.Println(C_A_1, C_A_2, C_A_3, C_A_4, C_A_5, C_A_6, C_A_7, C_A_8)
}
