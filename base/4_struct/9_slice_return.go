package main

import "fmt"

// 检验结构体没的 slice 在返回后是否还和结构体有关系
// 结论，会改变底层 array

type StructSlice struct {
	data []byte
}

func (s *StructSlice) Get() []byte {
	return s.data[2:3]
}

func main() {
	s := StructSlice{[]byte("1234567890")}

	s1 := s.Get()
	s1[0] = 'a'

	fmt.Printf("%s\n", s)
}
