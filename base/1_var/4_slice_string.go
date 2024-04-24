package main

import "fmt"

type Nodes []string

// Nodes转换为字符串
func (n *Nodes) String() string {
	return fmt.Sprintf("%s", n)
}

func main() {
	nodes := Nodes{"1", "2", "3"}
	fmt.Println(nodes)
}
