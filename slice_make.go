package main

import "fmt"

func main() {
	//var slice2 []int = make([]int, 10)
	slice1 := make([]int, 5, 10)
	fmt.Printf("%v\n", slice1)
	//slice1[0] = 100
	slice1 = slice1[:cap(slice1)]
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	fmt.Printf("%v\n", slice1)
	fmt.Printf("slice1 length:%d, cap:%d\n", len(slice1), cap(slice1))

	var slice2 []int
	fmt.Printf("%#v, %t\n", slice2, slice2 == nil)

	slice2 = make([]int, 5, 10)
	//slice2 = slice2[:cap(slice2)]
	fmt.Printf("%#v, %t\n", slice2, slice2 == nil)
	for i, v := range slice2 {
		slice2[i] = 5 * i
		fmt.Printf("%v === %v\n", i, v)
	}
	fmt.Printf("%#v\n", slice2)

	slice3 := []int{1, 3, 4}
	fmt.Println(slice3, len(slice3), cap(slice3))

	var slice4 []int  // nil
	// slice4[0] = 10	  // 直接赋值会出错，没有分配内存空间
	slice4 = append(slice4, 10, 20, 30)	// 不mark，直接append
	fmt.Println(slice4)
}