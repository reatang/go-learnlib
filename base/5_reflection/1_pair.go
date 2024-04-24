package main

import (
	"fmt"
	"io"
	"os"
)

/*
变量了内在结构对（pair）
*/

func main() {

	// pair<type:string, value:"123">
	a := "123"
	fmt.Println(a)

	var aInterface interface{}
	aInterface = a

	a1, ok := aInterface.(string)
	if ok {
		fmt.Println(a, a1)
	}

	// 转换为别的接口类型
	f, err := os.OpenFile("/tmp/golang_learn.out", os.O_RDWR, 0)
	if err == nil {

		// 类型转换
		w := io.Writer(f)
		w.Write([]byte("123123123\n"))
	} else {
		fmt.Println(err)
	}
}
