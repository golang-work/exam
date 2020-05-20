package main

import (
	"fmt"
	"strings"
)

func main(){
	// 字符串比较
	str := []byte("我")
	fmt.Println(len(str))
	fmt.Println(string(str) == "我")

	fmt.Println(strings.EqualFold("go", "GO"))

	// 是否存在某个字符或子串
	// chars 中任何一个 Unicode 代码点在 s 中，返回 true
	fmt.Println(strings.ContainsAny("this is my go lang", "is")) // true
	fmt.Println(strings.ContainsAny("this is my go lang", "siq")) // true	//

	// 子串 substr 在 s 中，返回 true
	fmt.Println(strings.Contains("this is my go lang", "si")) // false
	fmt.Println(strings.Contains("this is my go lang", "is")) // true

	// Unicode 代码点 r 在 s 中，返回 true
	fmt.Println(strings.ContainsRune("this is my go lang", rune('我'))) // false
	fmt.Println(strings.ContainsRune("this is 我 go lang", rune('我'))) // true
}
