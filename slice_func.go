package main

import "fmt"

func main() {
	// slice是引用类型
	slice1 := []int{1,2,3,4,5}
	updateS(slice1)
	fmt.Printf("%#v", slice1)
}

func updateS(slice []int) {
	for key, value := range slice {
		slice[key] = value * 2
	}
}