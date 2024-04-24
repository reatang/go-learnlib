package main

import (
	"fmt"
	"unsafe"
)

// byte slice 和 string 的转换优化

// byte slice 转 string
func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string 转 byte slice
func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// byte slice 转 string
func SliceByteToString2(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// string 转 byte slice
func StringToSliceByte2(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// 大写转小写
func lowerASCII(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func main() {

	// 老的写法
	b1 := []byte{'a', 'b', 'c', 'd'}
	s1 := SliceByteToString(b1)
	fmt.Println(s1)

	s2 := "abcd"
	b2 := StringToSliceByte(s2)
	fmt.Println(b2)

	// 1.20的写法
	s3 := SliceByteToString2(b1)
	fmt.Println(s3)

	b3 := StringToSliceByte2(s2)
	fmt.Println(b3)

	// 大写转小写
	var b4 byte = 'A'
	fmt.Println(b4, lowerASCII(b4))
}
