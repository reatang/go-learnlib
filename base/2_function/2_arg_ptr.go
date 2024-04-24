package main

import "fmt"

/*
	函数的指针参数

	本示例描述的时候，调用参数时，指针的实参和形参是不同的内存地址指向共同的值
	这和 C 是不一样的。
	结论：
	所以golang中无法通过改变形参指针存储的地址，来改变实参存储的地址。

	应用：
	但是可以传入实参的地址，相当于双层指针，可以改变实参指向。
*/

func FuncPtrArg(pi *int) {
	fmt.Printf("形参p1的地址：%p，实参pi的内容%p\n", &pi, pi)
}

func FuncPtrArg2(ppi **int, pi2 *int) {
	*ppi = pi2
}

func main() {
	i := 1
	pi := &i

	fmt.Printf("i的地址：%p\n", &i)

	fmt.Printf("实参pi的地址：%p，实参pi的内容%p\n", &pi, pi)
	FuncPtrArg(pi)

	// 将pi指向i2
	i2 := 2
	FuncPtrArg2(&pi, &i2)
	fmt.Printf("实参pi的地址：%p，实参pi的内容%p，实参pi指向的值内容%d\n", &pi, pi, *pi)

	// 上面 FuncPtrArg2 函数，相当于
	pi = &i2
	fmt.Printf("实参pi的地址：%p，实参pi的内容%p，实参pi指向的值内容%d\n", &pi, pi, *pi)

	// 将 i2值改为 i
	*pi = i
	fmt.Printf("i2 = %d, i = %d\n", i2, i)
}
