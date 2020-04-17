package main

import (
	"./packme"
	"fmt"
)

func main() {
	p := new(packme.Person)
	// 报错，类型 Person 被明确的导出了，但是它的字段没有被导出
	//p.firstName = "qing"

	// 可以通过OOP里的getter 和 setter 方法实现
	p.SetFirstName("qing")
	fmt.Println(p.FirstName())
}