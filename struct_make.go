package main

import "fmt"

// 定义方式：一
//type product struct {
//	name, intro string
//}

// 定义方式：二
type product struct {
	name,
	intro string
	price float64
}


func main() {
	// 定义变量及赋值
	//p1 := product{"vip", "test pro", 10}

	// 只定义变量
	var p1 *product
	var p2 product
	//p1 = &product{intro:"aaaa"}
	//p1.name = "###"
	//fmt.Printf("p1 type:%T, p1: %#v\n", p1, p1)

	// 错误，指针不能赋值给字面变量
	//p2 = new(product{})

	p2.name = "php"
	p1 = &p2

	fmt.Printf("p1 type:%T, value:%v\n", p1, p1)
	fmt.Printf("p2 type:%T, value:%v\n", p2, p2)
}
