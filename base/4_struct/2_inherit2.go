package main

import "fmt"

/*
复杂的继承，来源于 go-redis
同样的操作对不同的类可以抽象出来，达到多态的目的
*/

type cmdable func(say string) error

func (c cmdable) CmdSay(say string) {
	// 可以执行一些通用的操作，或者组装参数
	fmt.Println("cmdable before")

	c(say)

	// 可以执行一些后置操作或者组装返回值
	fmt.Println("cmdable after")
}

type A struct {
	Name string
	Age  int

	cmdable
}

func (a *A) cmdA(say string) error {
	fmt.Printf("%s say: %s\n", a.Name, say)

	return nil
}

type B struct {
	cmdable
}

func (b *B) cmdB(say string) error {
	fmt.Printf("B say: %s\n", say)

	return nil
}

func main() {
	a := A{Name: "reatang", Age: 18}
	a.cmdable = a.cmdA
	a.CmdSay("Hello Cmd")

	b := B{}
	b.cmdable = b.cmdB
	b.CmdSay("What the Cmd？！")
}
