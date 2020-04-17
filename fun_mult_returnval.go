package main

import "fmt"

func main() {
	sumPrice := sumPrice(1, 2, 3)
	fmt.Println(sumPrice)
}

func sumPrice(i, j, k int) int {
	return  i + j + k
}