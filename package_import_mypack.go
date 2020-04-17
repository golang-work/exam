package main

import (
	// 当使用 _ 来做为包的别名时，只执行它的 init 函数并初始化其中的全局变量
	// 当使用 . 来做为包的别名时，使用包内成员不用加包名
	//. "./book1"
	_ "./book"
	"./book2"
	"fmt"
)

func main() {
	fmt.Println(book2.MakeStr())
}
