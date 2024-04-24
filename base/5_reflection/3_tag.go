package main

import (
	"fmt"
	"reflect"
)

/*
使用反射解析结构体Tag
本用例使用 对象指针 和 对象本身 两种方式
*/

type User struct {
	Name string `info:"name" doc:"这是一个名字"`
}

// 用指针取
func findTagPtr(v *User) {

	ptr := reflect.TypeOf(v) // 这里返回的是一个ptr类型的 Type信息
	t := ptr.Elem()          // 相当于取指针指向的元素 *XX

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("info")

		fmt.Println(tag)
	}
}

// 用本身取
func findTag(v User) {
	t := reflect.TypeOf(v) // 这里返回的是类型本身 User 的 Type信息

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("info")

		fmt.Println(tag)
	}
}

func main() {

	var user User

	findTag(user)
	findTagPtr(&user)
}
