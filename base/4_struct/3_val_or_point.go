package main

import "fmt"

// 对于返回的接口是值还是指针的区别
/*
	结论：
	对于接口，【值】和【指针】属于两个不同的类型
	结构体只会有其中一种实现了接口
*/

var _ v3_interface = &v3_1{} // 此写法用来判断结构体是否实现了某接口

type v3_1 struct {
	A string
}

type v3_interface interface {
	S(name string)
	V3() string
}

// 只有 v3_1的指针类型实现了 v3_interface 接口
func (v *v3_1) S(name string) {
	v.A = name
}
func (v *v3_1) V3() string {
	return v.A
}

func getV3() v3_interface {
	return &v3_1{"defalut"}
}

func setV3(v3 v3_interface) {
	v3.S("setV3")
}

func main() {
	v1 := getV3()
	setV3(v1)
	fmt.Println(v1.V3())
}
