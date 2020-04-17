package main

import "fmt"

func main() {
	x := min(29, 3, 34, 123, 2)
	fmt.Printf("the min value is: %d\n", x)
	slice := []int{34, 545, 76, 33}
	x = min(slice...)
	fmt.Printf("the min value slice is: %d\n", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, value := range s {
		if value < min {
			min = value
		}
	}

	return  min
}