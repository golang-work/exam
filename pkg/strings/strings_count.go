package main

import (
	"fmt"
	"strings"
)

func main(){
	// 字符串比较
	fmt.Println(strings.Count("google.com", "o"))	// 3
	fmt.Println(strings.Count("google.com", ""))	// 11
	fmt.Println(strings.Count("谷歌中国", ""))	// 5

	fmt.Printf("%q\n", strings.Fields("  aa b	b         cc"))
	fmt.Printf("%q\n", strings.Split("  aa b	b         cc", "b"))
	fmt.Printf("%q\n", strings.SplitN("  aa b	b         cc", "b", 2))
	fmt.Printf("%t\n", strings.HasPrefix("baidu.com", "baidu1"))
	fmt.Printf("%d\n", strings.Index("baidu.com", "i"))
}
