package main

import "fmt"

// 合并两个切片
func copyMerge() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	size := len(a) + len(b)
	merge := make([]int, size)
	copy(merge, a)
	copy(merge[len(a):], b)

	showSlice(a, b, merge)
}

// 将源内容添加到目标的某个位置
func copyReplace() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}

	// 从 0 copy 到 len(a)
	copy(b, a)
	showSlice(a, b)

	// 从 len(a) copy 1个
	copy(b[len(a):len(a)+1], a)
	showSlice(a, b)

	// 从 len(a) copy 到最后
	copy(b[len(a):], a)
	showSlice(a, b)
}

// 复制到新的切片总
func copyNew() {
	a := []int{1, 2, 3}
	b := make([]int, len(a)) // 目标的长度一定要 >= 源数据的长度

	copy(b, a)

	showSlice(a, b)
}

func main() {
	// copyNew()
	// copyReplace()
	copyMerge()
}

func showSlice(l ...[]int) {
	for _, s := range l {
		fmt.Printf("len:%d, cap:%d, v:%v, ptr:%p\n", len(s), cap(s), s, s)
	}
}
