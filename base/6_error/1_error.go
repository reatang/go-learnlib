package main

import (
	"errors"
	"fmt"
)

/*
	1、Sentinel Error：自定义错误类型。尽量不要定义自己的error类型，最好使用标准库的。
	2、Error Types：自定义Error接口的实现。可以实现标准库的Error()接口来定义error，但是也要尽量避免。
	3、Opaque errors：1、不处理错误，直接返回。2、暴露接口而不暴露类型。
*/

func ReturnError() error {
	return errors.New("有错误了")
}

func main() {
	err := ReturnError()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("完成")
}
