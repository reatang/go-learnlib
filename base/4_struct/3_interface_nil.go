package main

// 接口的零值不是nil
// 接口的零值是一个包含nil指针的结构体
type Person interface {
	Say() string
}

type StructUserA struct {
	Name string
}

func (t *StructUserA) Say() string {
	return "你好"
}

func setNil() *StructUserA {
	return nil
}

func main() {
	var p Person
	p = setNil()
	if p == nil {
		println("p is nil")
	} else {
		println("p is not nil")
	}
	// p is not nil

	// 如何判断接口的零值是不是nil
	// 1. 判断接口的类型
	if p == nil || p.(*StructUserA) == nil {
		println("p is nil")
	} else {
		println("p is not nil")
	}
	// 2. 判断接口的值

}
