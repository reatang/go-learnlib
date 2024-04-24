package main

import (
	"fmt"
	"github.com/pkg/errors"
	err_lib "errors"
)

/*
	第三方包：github.com/pkg/errors
	包装错误，优化表达
*/

/*
	1、定义一个简单的包装
*/
type BusinessError struct {
	Biz string
	Err error
}

func handle() error {
	e := &BusinessError{"some Biz", nil}

	if e.Biz == "some Biz" {
		e.Err = errors.New("Some Error")
	}

	if e.Biz == "456" {
		e.Err = errors.New("Some Error 2")
	}

	return e.Err
}

/*
	2、第三方库包装
	在自己写三方库，或者一些多个项目共同使用的库的时候，依然不建议使用wrap包装错误，而是直接返回错误的根因（root error）
	只有在应用程序写业务的时候，因为需要错误堆栈信息，才需要包装错误栈。
*/
func e1() error {
	return errors.New("error 1")
}

func e2() error {
	err := e1()
	if err != nil {
		return errors.Wrap(err, "error 2")
	}
	return nil
}

func e3() error {
	err := e1()
	if err != nil {
		return err
	}
	return nil
}

/*
	3、1.13 更新的 Unwrap 接口，统一了带错误的结构体解包Sentinel Error的标准
*/
type QueryError struct {
	Query string
	Err error
}

func (q QueryError) Error() string {
	return q.Query
}

func (q *QueryError) Unwrap() error {
	return q.Err
}

func handle2() error {
	e := &QueryError{"some Biz", nil}
	e.Err = err_lib.New("原始错误")
	
	return e
}




func main() {
	// 1、
	err := handle()
	if err != nil {
		//fmt.Println(err)
	}
	
	// 2、
	err2 := e3()
	if err2 != nil {
		//fmt.Printf("Fatel: %s\n", errors.Cause(err2))
		//fmt.Printf("Stack: %+v\n", err2)
	}
	
	// 3、
	err3 := handle2()
	var E1 *QueryError
	if err_lib.As(err3, &E1) { // AS断言错误类型
		fmt.Printf("%s", E1.Query)
	}
}
