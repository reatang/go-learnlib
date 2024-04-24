package main

import "fmt"

// 实现类似class功能


// 1、重写方法，并调用父方法
type AClass struct {
	name string
}

func (a AClass) Name() string {
	return a.name
}

func (a AClass) Name2() string {
	return fmt.Sprintf("name2: %s", a.name)
}

type BClass struct {
	AClass // 继承
}

func (b BClass) Name() string {
	return fmt.Sprintf("我叫%s", b.AClass.Name())
}

// 2、结构体继承接口抽象
type AInterface interface {
	SayHello() string
}

type AA struct {
	SayWord string
}

// AA 和 *AA 是两个不同的实现，所以只有*AA实现了 AInterface
func (a *AA) SayHello() string {
	return a.SayWord
}

type BB struct {
	AInterface // 相当于 filed名称为 AInterface，类型是 AInterface
}

func (b BB) SayHello() string {
	return b.AInterface.SayHello()
}

func NeedAInterface(a AInterface)  {
	fmt.Println(a.SayHello())
}

func main() {
	
	// 1、结构体重写
	b := BClass{AClass{"reatang"}}
	fmt.Println(b.Name())
	fmt.Println(b.Name2()) // 直接使用父方法
	
	// 2、接口依赖重写
	bb := BB{&AA{"wow"}}
	// 此处是BB实现了接口，所以只能传BB而不能传*BB
	NeedAInterface(bb)
}
