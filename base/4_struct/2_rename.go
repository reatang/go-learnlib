package main

import (
	"fmt"
	"go-learn/base/base/4_struct/rename"
)

/*
	转换类型并操作，是否影响原变量研究

	结论1：
	别名有两种形式，这算是go语言中比较细节的语法了
	1、定义新数据类型
		定义方法：type 自定义数据类型 底层数据类型
		属于不同类型，混用需要类型转换
	2、类型别名
		定义方法：type 数据类型的别名 = 底层数据类型
		属于相同类型，混用无需类型转换

	结论2：
	1、本包内可以自由的做如上两种操作。
	2、跨包最好只使用 自定义数据类型，别名主要用于导出内部结构体给外部用
	3、最常用的还是对标量做 自定义数据类型，为标量添加方法以实现某些接口
	4、对指针做 自定义数据类型 估计只有两种情况，slice 和 map，因为这两个类型表象是普通类型，底层是指针类型
*/

type OriginA struct {
	A string
}

// 这种应用场景比较多，主要用于定义新类型并添加新的方法或者接口实现
type RenameA OriginA

// 值转换
func value() {
	oa := OriginA{"Origin"}
	ra := RenameA(oa)

	ra.A = "Rename"

	// 结论，是完全的两个变量，无关系
	fmt.Println(oa, ra)
}

// ------------------------------------------------------

// 这种应用场景很小，没啥用还，代码可读性也不好
type RenameB *OriginA

/*
// 指针类型不能直接添加接收器
func (b RenameB) Change(newName string) {
	b.A = newName
}
*/

type RenameC = *OriginA

func (c RenameC) Change(newName string) {
	c.A = newName
}

// 指针转换
func point() {
	oa := &OriginA{"Origin"}

	// 值别名，不能转换指针
	// ra := RenameA(oa)
	rb := RenameC(oa)

	// 直接赋值，可以改变
	rb.A = "Rename1"
	fmt.Println(oa, rb)

	// 方法调用，可以改变
	rb.Change("Rename2")
	fmt.Println(oa, rb)
}

// ------------------------------------------------------

// 最好不要用，没啥用（主要用于导出internal中的类型给外部使用）
type RenameD = rename.InnerA

/*
// 这样会报错
func (d RenameD) Change(name string) {
	d.A = name
}
*/

type RenameE rename.InnerA

func (e *RenameE) Change(name string) {
	e.A = name
}

func renameAddMethod() {
	// 1、在相同的包下，定义了别名，原类型也可添加方法，所以有点多此一举
	oa := OriginA{"Origin"}
	oa.Change("我竟然是A的方法")
	fmt.Println(oa)

	// 2、跨包支持这种吗
	// 结论，不支持

	// 3、跨包只能新建类型，而且因为是值转换，所以是新的变量
	// tips：属性相同的两个结构体，竟然可以自由转换！！
	oe := RenameE(oa)
	_ia := rename.InnerA(oa) // 不管是别名还是本体，都能自由转换
	oe.Change("EEE")
	fmt.Println(oa, oe, _ia.MyName())

	// 4、新类型可以用老类型的方法吗
	// 结论，当然是不可以的
	ia := rename.InnerA{"ia"}
	ie := RenameE(ia)
	ia.MyName()
	ie.Change("123")
	fmt.Println(ia, ie)
}

func main() {
	// value()
	// point()
	renameAddMethod()
}
