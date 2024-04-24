package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 匿名字段的反射和序列化过程
// 如果匿名字段是一个结构体，那么在序列化时，会把匿名字段的所有字段都会序列化出来
// 如果匿名字段是一个指针，那么在序列化时，会把匿名字段的所有字段都会序列化出来

type AAA struct {
	AName string `json:"a_name"`
}

type BBB struct {
	BName string `json:"b_name"`

	AAA
}

func main() {
	var b BBB
	typeB := reflect.TypeOf(b)
	for i := 0; i < typeB.NumField(); i++ {
		fmt.Printf("f: %s, %v\n", typeB.Field(i).Name, typeB.Field(i).Anonymous)
	}

	b2, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b2))

	s1 := "{\"b_name\":\"\",\"a_name\":\"123\"}"
	var b3 BBB
	err = json.Unmarshal([]byte(s1), &b3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", b3)
}
