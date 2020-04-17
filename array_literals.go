package main

import "fmt"

func main() {
	// 声明长度及类型的同时给定值，其它未给的将使用零值
	//arrayAge := [5]int{10, 20}
	//arrayAge := [5]int{}

	// 根据给定的值来决定数组的长度
	//arrayAge := [...]int{10, 20}

	// 用切片的方式定义数组
	//arrayAge := []int{20, 30}

	// 定义数组时给指定的key赋值
	arrayAge := [...]int{3:10, 8:20}
	arrayAge2 := [...]int{3:10, 8:20}
	//arrayAge := []int{3:10, 8:20}

	//var arrayAge [3]int = {1,2}
	//var (
	//	i1 = 10
	//	arrayAge [4]int
	//)
	//arrayAge[1] = 100

	fmt.Printf("arrayAge type:%T, len:%d, values:%#v\n", arrayAge, len(arrayAge), arrayAge)
	fmt.Printf("arrayAge ptr:%v\n", &arrayAge)
	for key, value := range arrayAge {
		fmt.Printf("arrayAge index:%d value:%d\n", key, value)
	}
	//fmt.Printf("i1 value %d\n", i1)

	fmt.Printf("arrayAge == arrayAge2 : %t", arrayAge == arrayAge2)
}