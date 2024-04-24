package main

import (
	"fmt"
	"reflect"
)

/*
利用pair 反射
本用例使用的是对象本身，而不是对象指针
*/

type RefUser struct {
	Name string
	Age  int
}

func (t RefUser) Say() {
	fmt.Println(t.Name, "say Hello")
}

func reflectVar(v interface{}) {
	fmt.Printf("type: %s \n", reflect.TypeOf(v))
	fmt.Printf("value: %s \n", reflect.ValueOf(v))
}

func reflectStruceField(v interface{}) {

	// 类型信息
	typeInfo := reflect.TypeOf(v)

	// 类型内容
	valueInfo := reflect.ValueOf(v)

	// 类型名称 相当于数据类型的名字，种类 go的基础数据类型
	fmt.Printf("类型名称：%s, 种类：%s\n", typeInfo.Name(), typeInfo.Kind())

	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		value := valueInfo.Field(i).Interface()

		fmt.Printf("name:%s, type:%s, value:%v \n", field.Name, field.Type, value)
	}

	// 1、直接使用方法名称调用方法
	valueInfo.MethodByName("Say").Call(nil)

	// 2、获取方法原型
	for i := 0; i < typeInfo.NumMethod(); i++ {
		method := typeInfo.Method(i)

		fmt.Printf("name:%s, type:%s\n", method.Name, method.Type)

		// 该处为独立func变量，需要传入上下文
		method.Func.Call([]reflect.Value{valueInfo})
	}
}

func main() {

	s1 := "qweqwe"
	reflectVar(s1)
	fmt.Println("------------")
	user := RefUser{"reatang", 12}
	reflectStruceField(user)
}
