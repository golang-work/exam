package main

import "fmt"

func main() {
	//slice1 := [3]int{1,2,3}[:]
	//slice1 := []int{1, 2, 3}
	//slice1 = slice1[:cap(slice1) + 1]
	//var slice1 = []int{}
	//fmt.Printf("slice cpa:%d\n", cap(slice1))
	//fmt.Printf("slice length:%d\n", len(slice1))
	//fmt.Printf("slice1 type:%T, value %v\n", slice1, slice1)

	arr := [10]int{1, 2, 3, 4, 5}
	slice := arr[:]

	fmt.Printf("%v  %v\n", slice[:5], slice[5:])

	//slice = slice[:cap(slice)]

	fmt.Println(arr)
	fmt.Printf("slice length:%d, cap:%d, %v", len(slice), cap(slice), slice)
}