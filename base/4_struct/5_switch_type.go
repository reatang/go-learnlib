package main

import "fmt"

// 类型隐式转换

type I1 interface {
	Say() string
}

type I2 interface {
	Walk()
}

type I3 interface {
	Jump()
}

type A struct {
	name string
}

func (t A) Say() string {
	return fmt.Sprintf("%s say: %s", t.name, "Hello")
}

func (t A) Walk() {
	fmt.Println(t.name, "在走")
}

// 多重断言
func action(i interface{}) {

	switch v := i.(type) {
	case I2:
		v.Walk()
	case I3:
		v.Jump()
	}
}

func main() {
	a1 := A{name: "reatang"}

	action(a1)
}
