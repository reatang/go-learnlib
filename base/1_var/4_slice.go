package main

import "fmt"

/**
切片细节

切片的数据结构
type SliceHeader struct {
	Data uintptr // 底层数组的指针
	Len  int // 切片长度
	Cap  int // 切片容量
}

*/

func main() {
	slice0()

	// slice1()

	// slice2()
}

func slice0() {
	var s1 []string

	s1 = append(s1, "123")

	fmt.Printf("%v\n", s1)

	type S1 struct {
		s1 []byte
	}

	ss1 := S1{}
	if ss1.s1 == nil {
		ss1.s1 = append(ss1.s1, 'a')
		fmt.Println("is nil")
	}
	fmt.Printf("%s\n", ss1.s1)

	var ss2 []string
	if ss2 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Printf("%+v\n", ss1.s1)
	}
}

func slice1() {
	// 1、append超过cap，切片的反应
	s1 := make([]string, 2, 3)
	s1[0] = "1"
	showSlice(s1)
	s2 := append(s1, "2") // 此处不是向1号索引插入，而是从Len长度后面插入,已超过s1管理范围
	s3 := append(s2, "3") // 此处再插入，则超过Cap值，已超过底层数组长度
	s1[1] = "X"
	s2[2] = "x"
	showSlice(s1)
	showSlice(s2)
	showSlice(s3)
	s3[3] = "4"

	s3[1] = "-2" // 所以，对s3的修改不会影响s1、s2。因为底层数组重新申请了内存
	showSlice(s3)
	// 结论，len定义的是切片的长度，而cap定义的是底层数组的预定长度
}

// 自动裁剪（看起来，没什么卵用）
func slice2() {
	data := make([]string, 0)

	data = append(data, "1")
	data = append(data, "2")
	data = append(data, "3")
	data = append(data, "4")
	data = append(data, "5")

	data2 := data[0:len(data)]

	showSlice(data)
	showSlice(data2)
}

func showSlice(s []string) {
	fmt.Printf("len:%d, cap:%d, v:%v, ptr:%p\n", len(s), cap(s), s, s)
}
