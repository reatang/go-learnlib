package main

import "fmt"

/*
close	用于在管道通信中关闭一个管道
len、cap	len 用于返回某个类型的长度（字符串、数组、切片、字典和管道），cap 则是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道）
new、make	new 和 make 均用于分配内存，new 用于值类型和用户自定义的类型（类），make 用于内置引用类型（切片、字典和管道）。它们在使用时将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针，也可以用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，所以 make 不仅分配内存地址还会初始化对应类型。
copy、append	分别用于切片的复制和动态添加元素
panic、recover	两者均用于错误处理机制
print、println	打印函数，在实际开发中建议使用 fmt 包
complex、real、imag	用于复数类型的创建和操作
*/

func runArrayAndSlice() {
	// array、slice
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr1 len = %d, arr1 cap = %d, arr1 = %v\n", len(arr1), cap(arr1), arr1)
	sli1 := arr1[1:2]
	fmt.Printf("sli1 len = %d, sli1 cap = %d, sli1 = %v\n", len(sli1), cap(sli1), sli1)

	sli1 = append(sli1, 7, 8, 9, 10, 11)
	fmt.Printf("sli1 len = %d, sli1 cap = %d, sli1 = %v\n", len(sli1), cap(sli1), sli1)

	sli2 := []int{0, 0, 0}
	copy(sli2, sli1)
	fmt.Printf("sli2 len = %d, sli2 cap = %d, sli2 = %v\n", len(sli2), cap(sli2), sli2)
}

type A1 struct {
	A int
}

type A2 struct {
	A *int
}

// 安全的使用对象的值传给另一个对象
func runNewVarPtr() {
	var a = A1{A: 1}

	// 不安的传值
	a2 := A2{}
	a2.A = &a.A

	*a2.A = 10
	fmt.Println(a.A) // 会影响原对象

	// 安全的传值
	i := a.A
	a3 := A2{}
	a3.A = &i
	*a3.A = 20
	fmt.Println(a.A, i)

	// 简化法 安全传值
	a4 := A2{}
	a4.A = IntPtr(a.A)
	*a4.A = 30
	fmt.Println(a.A, *a4.A)
}

// 返回新变量的地址
func IntPtr(i int) *int {
	return &i
}

func main() {
	// runArrayAndSlice()

	runNewVarPtr()
}
