package main

import "fmt"

func main() {
	for i :=0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("i value:%d\n", i)
	}

	for j :=0; j < 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Printf("j value:%d\n", j)
	}
}
