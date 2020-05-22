package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "qing我"
	fmt.Printf("%s\n", strings.TrimSpace(strings.Repeat(name, 10)))

	// len 函数返回的是字节长度
	// 如果希望返回字符长度可以借助utf8包里的方法
	fmt.Println(len(name)) // 7
	fmt.Println(utf8.RuneCountInString(name)) // 5

	name = "我的"
	fmt.Println(len(name)) // 6
	fmt.Println(utf8.RuneCountInString(name)) // 2
}