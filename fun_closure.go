package main

import "fmt"

var x int

func main() {
	// 闭包函数保存并积累其中的变量的值，
	// 不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
	a2 := adder2()
	fmt.Println(a2(10))
	fmt.Println(a2(20))
	fmt.Println(a2(30))
	fmt.Printf("%#v", a2)

	//fmt.Println(adder2()(10))
	//fmt.Println(adder2()(20))
	//fmt.Println(adder2()(30))
}

func adder2() func(b int) int{
	return func(b int) int{
		x += b
		return  x
	}
}