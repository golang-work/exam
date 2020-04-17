package main

import "fmt"

func main() {
	s1From := []int{1, 2, 3}
	s1To := make([]int, 10)
	n := copy(s1To, s1From)
	fmt.Println(s1To)
	fmt.Printf("copied %d elements\n", n)

	//sl3 := []int{1, 2, 3}
	sl3 := make([]int, 0)
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}