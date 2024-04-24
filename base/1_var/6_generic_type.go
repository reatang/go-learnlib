package main

import "fmt"

// 泛型

type G1[T int | float32 | float64] struct {
	A T
	B T
}

func (t *G1[T]) Add() T {
	return t.A + t.B
}

func main() {
	g := G1[int]{A: 1, B: 2}
	fmt.Println(g.Add())
}
