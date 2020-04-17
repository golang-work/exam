package main

import "fmt"

// 普通定义
var name string = "qing"

// 批量定义，类型可以不显示定义，使用类型推导
var address, email = "shang hai", "490702087@qq.com"
var (
	age = 100
	score = 98.88
)

func main() {
	// 简短形式，注意，这种方式只能在函数体内定义，全局不能这样定义
	price := 100.78

	fmt.Printf("my name is:%s, address:%s, email:%s", name, address, email)
	println()
	fmt.Printf("my age is:%d, score:%f", age, score)
	println()
	fmt.Printf("price:%.2f", price)
	println()

	// 可以先批量定义（在只定义不赋值的情况下类型是必须给的），然后再批量赋值
	var a, b, c int
	a, b, c = 10, 20, 30
	fmt.Printf("price:%d", a, b, c)
	println()
}