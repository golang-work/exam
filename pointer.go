package main

import (
	"fmt"
)

const price  = 2

func main() {
	// 指针的零值是nil，对一个空指针的反向引用是不合法的，并且会使程序崩溃
	//var nameP *string
	//*nameP = 100
	//fmt.Printf("nameP value:%v, nameP value:%+v, nameP value:%#v\n", nameP, nameP, nameP)

	// 对于常量是不能取地址的
	//fmt.Printf("price pointer:%p\n", &price)
	count := 100
	fmt.Printf("count type:%T\n", count)
	fmt.Printf("count value:%d, pointer:%p\n", count, &count)

	countP := &count
	*countP = 102
	fmt.Printf("count value:%d, pointer:%p\n", count, &count)

	if  count == *(&count){
		fmt.Println("if ok...")
	}

}