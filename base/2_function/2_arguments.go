package main

import "fmt"

/*
函数的参数
*/

// 强类型参数
func F1(a int) {
	fmt.Printf("a=%d\n", a)
}

// 指针参数
type S1 struct {
	Id int
}

func F2(s *S1) {
	s.Id = 123456
}

func F2_(s S1) {
	s.Id = 654321
}

// 可变参数
func F3(lists ...int) {

	for _, v := range lists {
		fmt.Println("i=", v)
	}
}

func main() {
	F1(1)

	s1 := S1{1}
	F2(&s1)
	fmt.Println(s1.Id)
	F2_(s1)
	fmt.Println(s1.Id)

	// 可变形参
	F3(1, 2, 3, 4, 5, 6)
	// 可变实参
	F3([]int{6, 5, 4, 3, 2, 1}...)
}
