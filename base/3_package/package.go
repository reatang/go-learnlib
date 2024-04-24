package main

import (
	. "fmt" // 表示可匿名的导入，这样在使用过程中该包内方法无需包名

	_ "go-learn/base/base/3_package/lib"         // 特殊别名 表示只运行包内init()
	"go-learn/base/base/3_package/lib2"          // 全名
	xxx "go-learn/base/base/3_package/lib3/beta" // 包名和目录名称不一致
)

func main() {
	Println(lib2.LibTest2())
	Println(xxx.XXX())
}
