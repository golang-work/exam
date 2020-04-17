package main

import "fmt"

type struct1 struct {
	i1 int
	f1 float64
	str string
}

func main() {
	// 普通定义
	// var ms *struct1 = new(struct1)

	// 简洁定义
	//ms := &struct1{10, 15.5, "chris"}

	// 常规定义
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "chris"

	fmt.Printf("ms type:%T, value:%v\n", ms, ms)
}
