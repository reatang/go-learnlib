package main

import (
	"fmt"
)

/*
array 是值传递
slice 是引用传递

学习切片相关方法
len/cap
append/copy

*/

func s1() {
	// 数组 是固定长度的
	arr1 := [10]int{1, 2, 3, 4}
	fmt.Println(arr1)

	arr2 := [...]int{5, 6, 7, 8}
	fmt.Println(arr2)

	// 切片 切片是引用传递
	// 1、先声明，再分配空间
	var sliceA []int
	sliceA = make([]int, 10)
	fmt.Println(sliceA)

	// 2、直接创建切片
	sliceB := []int{11, 22, 33, 44}
	fmt.Println(sliceB)

	// 3、从已存在的数组创建（截取）切片
	slice1 := arr1[:]
	fmt.Println(slice1)

	slice2 := arr1[1:2:10] // [low_index, high_index, max_index] 从low取到high(不包含high)并预留到max_index
	fmt.Println(slice2)

	fmt.Println("----------")

	// 切片容量 切片的容量会按策略自动增加（在短的时候按2倍增加，长度到1024后按1.25倍增加）
	fmt.Println(len(slice2), cap(slice2)) // 长度：切片中实际的内容数量，容量 = max_index - low_index

	slice2 = append(slice2, 2222)
	fmt.Println(len(slice2), cap(slice2))

	fmt.Println("edit ----------")

	// 修改切片内容
	slice2[0] = 666
	fmt.Println(slice2, arr1, slice1) // 修改其中一个切片，底层数组和相关切片都会变化

	slice3 := append(slice2, 'B', 'C') // 此处需要注意，如果append的内容大于cap长度，会创建新切片，不会改变关联的数组和其他切片。
	fmt.Println(slice3, arr1)
	slice3[0] = 0
	fmt.Println(slice3, arr1, slice1, slice2) // append对切片容量范围的数据，也会修改

	fmt.Println("copy ----------")

	// copy 切片
	var slice_copy = make([]int, 5)
	slice_copy[4] = 1999
	copy(slice_copy, sliceB) // 复制会将目标切片对应长度的数据都覆盖
	slice_copy[0] = 11111    // 修改不会影响原来的切片
	fmt.Println(slice_copy, sliceB)

	fmt.Println("append ----------")

	arr3 := [5]int{1, 2, 3, 4, 5}
	slice4 := arr3[0:5:5]
	slice5 := append(slice4, 6) // 超过cap，新建切片，和老数据脱离联系
	slice4[1] = 200             // 未脱离关联的切片会影响老数据
	slice5[0] = 11              // 已脱离关联，老数据不会变化
	fmt.Println(arr3, slice4, slice5)
}

// 测试传入部分切片，修改后的结果
func s2() {
	a1 := []int{1, 2, 3, 4, 5, 6}

	s2_1(a1[1:2])

	// 事实证明，传入的新切片还是跟老的切片底层共享数组，相当于传入了部分指针
	fmt.Println(a1)
}
func s2_1(a1 []int) {
	a1[0] = 666
}

func main() {
	s1()
}
