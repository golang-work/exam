package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[: i + 1]
		slice1[i] = i
		fmt.Printf("the length of slice is %d\n", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}[4:]
	fmt.Printf("slice2 value:%#v, length:%d, cap:%d\n",
		slice2, len(slice2), cap(slice2))
	slice2 = slice2[2:4]
	fmt.Printf("slice2 value:%#v, length:%d, cap:%d\n",
		slice2, len(slice2), cap(slice2))
}