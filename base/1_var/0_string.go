package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str1 := "一个字符串"
	fmt.Println(str1)

	// here doc 格式
	multiStr1 := `第一行
wow
多行`

	fmt.Println(multiStr1)

	// 字符串长度 和 多字节字符串长度
	fmt.Println(len(multiStr1))                  // 20
	fmt.Println(len("你好"))                       // 6
	fmt.Println(utf8.RuneCountInString("你好"))    // 2
	fmt.Println(utf8.RuneCountInString("你好abc")) // 5

	// 取部分字符串
	fmt.Println("abcdefg"[1:4]) // bcd

	// 字符串连接
	fmt.Println("123" + "456")
	fmt.Println(fmt.Sprintf("%s%s", "123", "456"))

	// 字符串分割
	str2 := "1,2,3,4,5,6,7,8"
	fmt.Println(strings.Split(str2, ","))
	fmt.Println(strings.SplitN(str2, ",", 2))

	// 连接
	fmt.Println(strings.Join([]string{"1", "2", "3"}, ","))

	/**
	查找和替换
	*/
	// 包含
	fmt.Println(strings.Contains(str2, "4"))

	// 子字符串首次数显的索引位置
	fmt.Println(strings.Index(str2, ","))     // 从前开始数
	fmt.Println(strings.LastIndex(str2, ",")) // 从后开始数

	// 字符串修改，从 string -> []rune -> 修改 -> string
	str3 := "我爱中国"
	rune1 := []rune(str3)
	rune1[0] = '你'
	fmt.Println(string(rune1))

	// 前缀后缀匹配
	fmt.Println(strings.HasPrefix(str2, "1"))
	fmt.Println(strings.HasSuffix(str2, "8"))

	// 查找替换
	str4 := "你好，世界"
	replaceStr4 := strings.Replace(str4, "你好", "您好", 1)
	fmt.Println(replaceStr4)
}
