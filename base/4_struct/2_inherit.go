package main

import "fmt"

/*
结构体继承
*/

type StructA struct {
	Name string
	Age  int
}

func (t StructA) Asay() string {
	return fmt.Sprintf("name: %s", t.Name)
}

type StructB struct {
	StructA
	Color string
}

func main() {
	// 1、基本的创建结构体对象
	a := StructB{StructA{"reatang", 18}, "#FFFFFF"}
	fmt.Println(a)

	// 2、先创建，再赋值
	b := StructB{}
	b.Name = "reatang2"
	b.Age = 123
	b.Color = "#000000"
	fmt.Println(b.Asay())
}
