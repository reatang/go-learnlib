package main

import "fmt"

// DeleteSlice1 删除指定元素。截取法（修改原切片）
func DeleteSlice1(a []int, elem int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == elem {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

// DeleteSlice2 删除指定元素。拷贝法（不改原切片）
func DeleteSlice2(a []int, elem int) []int {
	tmp := make([]int, 0, len(a))
	for _, v := range a {
		if v != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

// DeleteSlice3 删除指定元素。 移位法1（修改原切片）
func DeleteSlice3(a []int, elem int) []int {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// DeleteSlice4 删除指定元素。移位法2（修改原切片）
func DeleteSlice4(a []int, elem int) []int {
	tgt := a[:0]
	for _, v := range a {
		if v != elem {
			tgt = append(tgt, v)
		}
	}
	return tgt
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	showSlice(s1)
	s1 = DeleteSlice4(s1, 5)
	showSlice(s1)
}

func showSlice(l ...[]int) {
	for _, s := range l {
		fmt.Printf("len:%d, cap:%d, v:%v, ptr:%p\n", len(s), cap(s), s, s)
	}
}
