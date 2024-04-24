package main

import (
	"fmt"
	"time"
)

func goError() {
	panic("soem error")
}

func goA() {
	defer func() { fmt.Println("defer 1") }()
	defer func() { fmt.Println("defer 2") }()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("有错误")
		}
	}()

	goError()
}

// defer依然是按栈方式运行
func goB() (ret string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("有错误")
			ret = "修改为 error的内容"
		}
	}()

	defer func() { fmt.Println("defer 1") }()
	defer func() { fmt.Println("defer 2") }()

	goError()
	ret = "success"

	return ret
}

func main() {
	go goA()
	ret := goB()

	fmt.Printf("修改返回值：%s\n", ret)

	time.Sleep(1 * time.Second)
}
