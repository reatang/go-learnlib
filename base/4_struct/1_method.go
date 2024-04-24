package main

import (
	"fmt"
	"go-learn/base/base/4_struct/method"
)

// 不能在别的包添加方法
// func (*method.AAA) SayB() string {
// 	return "say b"
// }

// 只能继承
type BBB struct {
	method.AAA
}

// 再添加新的方法
func (*BBB) SayB() string {
	return "say b"
}

func main() {
	a := BBB{}

	fmt.Println(a.SayB())
}
