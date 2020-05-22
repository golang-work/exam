package main

import "fmt"

type fun1 func(...interface{})

func main() {
	var aaa fun1	// 函数类型零值是nil
	fmt.Println(aaa)
	aaa = func(s ...interface{}) {
		fmt.Println(s...)
	}
	aaa("str", "ok", "...")

	fmt.Println(multiPly3Nums(1, 2,3))
}

func multiPly3Nums(a int, b int, c int) (product int) {
	return  a * b * c
}

//func multiPly3Nums(a int, b int, c int) (product int) {
//	product = a * b * c
//	return  product
//}

//func multiPly3Nums(a int, b int, c int) int {
//	var product int
//	product = a * b * c
//	return  product
//
//	//return a * b * c
//}


