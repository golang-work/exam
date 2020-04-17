package main

import "fmt"

type IntVector []int

func main() {
	fmt.Println(IntVector{1, 2, 3}.sum())
}

func (v IntVector) sum() (s int) {
	for _, x := range v {
		s += x
	}

	return
}
