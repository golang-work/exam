// 继承多个struct，有命名相同时
package main

import "fmt"

type A struct {
	a int
}

type B struct {
	a, b int
}

type C struct {
	A
	B
}

func main() {
	c := new(C)
	// 正确
	c.A.a = 100
	c.B.a = 200

	// 错误，程序没有办法来解决这种问题引起的二义性，必须由程序员自己修正
	//c.a = 300

	fmt.Println(c)
}