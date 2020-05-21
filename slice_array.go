package main

import "fmt"

func main(){
	var arr1 [3]int
	arr1[0] = 10
	fmt.Printf("arr1 type: %T, arr1 val: %#v\n", arr1, arr1)
	var arr2 = new([3]int)
	(*arr2)[0] = 20
	fmt.Printf("arr2 type: %T, arr2 val: %#v\n", arr2, arr2)
}
