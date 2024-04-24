package main

import "fmt"

/*
多态：接口实现
只要实现了接口的所有方法，则认为实现了某个接口
*/

type Person interface {
	Say() string
	Walk()
}

type StructUserA struct {
	Name string
	Age  int
}

func (t StructUserA) Say() string {
	return "你好"
}

func (t StructUserA) Walk() {
	fmt.Printf("%s正在走动", t.Name)
}

type StructUserB struct {
	Name string
	Age  int
}

func (t StructUserB) Say() string {
	return "沙雕"
}

func (t StructUserB) Walk() {
	fmt.Printf("%s正在飞行\n", t.Name)
}

//////
func Talk(p Person) {
	fmt.Println(p.Say())
	p.Walk()
}

func main() {
	a := &StructUserA{"aaa", 1}
	b := &StructUserB{"bbb", 1}

	Talk(a)
	Talk(b)
}
