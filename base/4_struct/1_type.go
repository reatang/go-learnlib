package main

import "fmt"

/*
定义结构体
初始化结构体

*/

// 定义结构体
type MyStruct struct {
	a int
}

type MyStruct2 struct {
	T1 *MyStruct
	T2 MyStruct
}

// 为结构体添加带对应接收器的函数
func (this MyStruct) say() string {
	return fmt.Sprintf("Hello %d", this.a)
}

func (this *MyStruct) setA(a int) {
	this.a = a
}

func main() {

	// 1、实例化对象然后赋值
	var myStruct1 MyStruct
	myStruct1.a = 123
	fmt.Printf("%p", &myStruct1)
	showStruct(&myStruct1)

	// 2、直接实例化对象
	var myStruct2 = MyStruct{a: 2}
	myStruct3 := MyStruct{a: 3}
	showStruct(&myStruct2)
	showStruct(&myStruct3)

	// 3、用new创建实例化后对象的指针
	myStruct4 := new(MyStruct)
	myStruct4.a = 4
	showStruct(myStruct4)

	// 结构体方法
	fmt.Println(myStruct1.say())
	myStruct1.setA(111)
	showStruct(&myStruct1)

	fmt.Println("----------")

	// 传值是深拷贝吗
	s2 := MyStruct2{T1: &MyStruct{a: 1}, T2: MyStruct{a: 2}}
	fmt.Printf("%v, %p\n", s2, &s2)
	showStruct2(s2)
	fmt.Printf("%v, %p", s2, &s2)
	// 答案，是深拷贝，但是不包括指针

	// (*response)(nil) 是什么意思，在net/http包中看到
	//someFunc := func(val interface{}) {
	//	fmt.Println(val)
	//}
	fmt.Println("\n-------------")
	s1 := string([]byte{'c', 'b', 'a'})
	m := (*MyStruct)(nil)
	var m2 *MyStruct
	fmt.Println(m, m2)
	fmt.Println(s1)
	//m.say()
	//m2.say()
	//someFunc(m)

}

func showStruct(s *MyStruct) {
	fmt.Printf("v:%v, ptr:%p\n", *s, s)
}

func showStruct2(s MyStruct2) {
	s.T1.a = 666
	s.T2.a = 777
	fmt.Printf("内部：%p\n", &s)
}
