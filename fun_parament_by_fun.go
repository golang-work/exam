package main

import "fmt"

func main() {
	callback(sum, 10, 20, 30, 2)
}

func sum(args ...int) int {
	result := 0
	for _, value := range args {
		result += value
		fmt.Printf("result add:%d\n", value)
	}
	return result
}

func callback(f func(...int) int, args ...int) {
	result := f(args...)
	fmt.Printf("result:%d\n", result)
}
