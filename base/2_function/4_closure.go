package main

import (
	"fmt"
	"time"
)

func main() {

	// 1、匿名函数
	f1 := func() {
		fmt.Println("定义了一个匿名函数")
	}

	// 执行匿名函数
	f1()

	// 2、定义并执行
	func() {
		fmt.Println("定义并立马执行")
	}()

	// 3、协程执行匿名
	go func() {
		fmt.Println("在协程中的匿名函数")
	}()

	time.Sleep(1 * time.Second)
}
