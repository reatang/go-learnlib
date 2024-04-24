package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 十分之一概率报异常
func f2() {
	time.Sleep(time.Second)

	i := rand.Intn(3) // 0、1、2

	if i == 2 {
		panic(errors.New("报错咯"))
	}

	log.Println("f2(): 正常执行中")
}

func f1() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("f1() defer2: recover " + err.(error).Error())
		}

		log.Println("f1() defer2: recover 执行完成")
	}()

	for {
		f2()
	}
}

func run() error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("run() defer2 recover:" + err.(error).Error())
		}
	}()

	go f1()

	defer func() {
		log.Println("run() defer1: run 执行完成")
	}()

	return nil
}

func main() {

	e := run()
	if e != nil {
		log.Fatal(e)
	}

	time.Sleep(10 * time.Second)
	log.Println("执行完成")
}
