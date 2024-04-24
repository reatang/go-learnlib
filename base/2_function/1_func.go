package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(A())
	fmt.Println(B(2))
	fmt.Println(C(1, 2))

	name, age := D("1", 2)
	fmt.Printf("My name is %s, age %d\n", name, age)

	fmt.Println(E("1", 2))
	fmt.Println(F(6, 6, 6, 6, 6))
	fmt.Println(G("wow"))

	h := H{"1"}
	fmt.Println(h.Hsay())
	h.SetName("2")
	fmt.Println(h.Hsay())
}

func A() int {
	return 1
}

func B(b int) int {
	return b * b
}

func C(a, b int) int {
	return a + b
}

func D(name string, age int) (string, int) {
	return name, age
}

func E(name string, age int) (say string) {
	say = fmt.Sprintf("My name is %s, age %d", name, age)
	return
}

func F(list ...int) string {
	var strList = make([]string, len(list))

	for index, value := range list {
		strList[index] = strconv.Itoa(value)
	}

	return strings.Join(strList, ", ")
}

func G(name string) string {

	defer func() {
		fmt.Println("我" + name + "还有话要说！")
	}()

	return name
}

type H struct {
	name string
}

func (this H) Hsay() string {
	return this.name
}

func (this *H) SetName(name string) bool {
	this.name = name

	return true
}
