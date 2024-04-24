package main

import "fmt"

// 函数指针绑定上下文运行

type ClassBind struct {
	A string
}

func (c *ClassBind) Say() {
	fmt.Println(c.A, "say: Hello!")
}

// type F func(*ClassBind)

func main() {

	a := &ClassBind{A: "reatang"}
	b := &ClassBind{A: "leeya"}

	a.Say()

	f := (*ClassBind).Say
	f(b)
}
