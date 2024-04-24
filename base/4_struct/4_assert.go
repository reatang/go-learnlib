package main

import "fmt"

/*
空接口与类型断言
*/

func varAssert(arg interface{}) {

	// 类型断言
	v, is := arg.(string)
	if is {
		fmt.Println("字符串：", v)
	}
}

func main() {

	s1 := "这是一个字符串"

	varAssert(s1)
}
