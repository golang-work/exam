package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[3:]
	fmt.Printf("arr1 type is:%T, value is:%#v\n", arr1, arr1)
	fmt.Printf("slice1 type is:%T, value is:%#v\n", slice1, slice1)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d \n", i, slice1[i])
	}
	fmt.Printf("the length of arr1 is %d\n", len(arr1))
	fmt.Printf("the length of slice1 is %d\n", len(slice1))
	fmt.Printf("the capacity of slice is %d\n", cap(slice1))

	slice1 = slice1[0:len(slice1) - 1]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
}