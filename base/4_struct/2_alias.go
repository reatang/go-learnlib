package main

import "fmt"

type MyInt int      // 定义新类型
type AliasInt = int // 定义别名

// 2、可以添加方法
func (i MyInt) Add(i2 int) int {
	return int(i) + i2
}

// 2、标量不能添加方法
// func (i AliasInt) Add(i2 int) int {
// 	return i + i2
// }

// 3、新类型可以添加方法
type MyInt2 = MyInt

func (i MyInt2) Sub(i2 int) int {
	return int(i) - i2
}

func main() {
	i1 := MyInt(1)
	i2 := AliasInt(1)

	fmt.Printf("%t, %t\n", i1, i2)

	// 1、不能直接判断相等
	// fmt.Printf("eq=%d", i1 == i2)
	fmt.Printf("eq=%t\n", int(i1) == i2)

	// 2、可以添加方法
	fmt.Println(i1.Add(2))

	// 3、别名添加方法，原类型也会添加方法
	// 但是跨包不能添加新方法
	i3 := MyInt2(3)
	fmt.Println(i3.Sub(1))
	fmt.Println(i1.Sub(1))

}
