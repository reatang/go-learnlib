package main

import (
	"context"
	"fmt"
)

// 空结构体配合context存储上下文数据

type Void struct{}
type Void2 struct{}

var UniqueKey1 Void
var UniqueKey2 Void2

func main() {
	ctx := context.WithValue(context.Background(), UniqueKey1, 123)

	// 拥有唯一的key
	i1 := ctx.Value(UniqueKey1)
	fmt.Println(i1)

	// 别的结构体也不能读取到
	i2 := ctx.Value(UniqueKey2)
	fmt.Println(i2)
}
