package main

import "fmt"

func main(){
	// Println 打印多个值时，值始终会用 " " 空格连接
	fmt.Println(1234, 1234)
	fmt.Println("abc", 1234)
	fmt.Println("abc", "123")
	fmt.Println()

	// Print 打印多个值时，如果都不是字符串则使用 " " 空格连接，否则直接打印在一起
	fmt.Print(1234, 1234)
	fmt.Println()
	fmt.Print("abc", 1234)
	fmt.Println()
	fmt.Print("abc", "1234")
	fmt.Println()
}
