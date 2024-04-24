package main

import "fmt"

type Config struct {
	a int
	b string
}

// 值覆盖，c 和 config 是两个新的指针
func (c *Config) Apply(config *Config) error {

	fmt.Printf("%p, %p\n", &c, c)
	fmt.Printf("%p, %p\n", &config, config)

	if config != c {
		// config = c // 无用操作
		*config = *c
	}
	return nil
}

// 值对值
func main() {
	c1 := &Config{1, "001"}
	c2 := &Config{2, "002"}

	// 值覆盖
	c1.Apply(c2)

	// 值覆盖
	// *c2 = *c1

	// 地址赋值，两个指针指向同一个地址
	// c2 = c1

	c2.a = 10

	// 指针自己的地址, 指向的地址, 指向的内容
	fmt.Printf("%p, %p, %v\n", &c1, c1, c1)
	fmt.Printf("%p, %p, %v\n", &c2, c2, c2)
}
