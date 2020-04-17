package main

import "fmt"


func main() {
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


