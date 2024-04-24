package main

import "testing"

// 测试 断言的性能 和 判断一个数字的性能是否一致
// 测试结论：断言比判断一个数字要慢20%，但是从设计来说，断言更符合语言多态特性

var colorType = 1
var Color2 = ColorA{"Green"}

type Colorable interface {
	Color(string) string
}

type ColorA struct {
	t string
}

func (c *ColorA) Color(s string) string {
	return c.t + s
}
func (c *ColorA) Color2222(s string) string {
	return c.t + s + "22222"
}

func ShowColor(colorable Colorable) {
	if a, ok := colorable.(*ColorA); ok {
		a.Color2222("reatang")
	}
}

func ShowColor2(typeId int) {
	if typeId == 1 {
		Color2.Color("reatang")
	}
}

// go test --bench=. --count=3 --benchmem --run=none --cpu=8 4_assert_performance_test.go
// go test --bench=. -benchtime=5s --benchmem --run=none --cpu=8 4_assert_performance_test.go

func BenchmarkAssert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ShowColor(&Color2)
	}
}

func BenchmarkAssertType(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ShowColor2(colorType)
	}
}
