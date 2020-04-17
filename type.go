package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

func main() {
	count := 100
	fmt.Printf("count type is:%T\n", count)
	fmt.Printf("bool type:%t\n", count - 100 > 10)
	for i :=0; i < 5; i++ {
		fmt.Printf("rand number:%d\n", rand.Int())
	}
	fmt.Printf("now time:%s\n", time.Now().Format("2006-01-02 15:04:05"))
	
	// 类型别名
	type TZ int
	var price1 TZ = 100
	price2 := 100
	var price3 int
	price3 = int(price1) + price2

	fmt.Printf("price type is:%T\n", price1)
	fmt.Printf("price3 value is:%d\n", price3)

	// string类型
	//var char = '你'
	var char = '\u00FF'
	fmt.Printf("char type is:%T, char value:%v\n", char, char)
	fmt.Printf("%c\n", char)

	fmt.Printf("%t\n", unicode.IsDigit('a'), unicode.IsDigit('1'))
}