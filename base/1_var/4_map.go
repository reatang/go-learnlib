package main

import "fmt"

/*
map 是引用传递

学习 delete 方法

*/
func main() {

	// 1、预分配长度的map
	var map1 map[string]string
	map1 = make(map[string]string, 2)
	map1["a"] = "a"
	map1["b"] = "b"
	map1["c"] = "c"
	fmt.Println(map1)

	// 2、不分配长度
	map2 := make(map[string]string)
	map2["A"] = "A"
	map2["B"] = "B"
	map2["C"] = "C"
	fmt.Println(map2)

	// 3、直接声明并赋值
	map3 := map[string]string{
		"AA": "AA",
		"BB": "BB",
		"CC": "CC",
	}
	map3["DD"] = "DD"
	fmt.Println(map3)

	// 遍历
	for k, v := range map3 {
		fmt.Printf("key:%s, value:%s\n", k, v)
	}

	// 删除
	delete(map3, "DD")

	fmt.Println(map3)
}
