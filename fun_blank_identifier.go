package main

import "fmt"

func main() {
	i1, _, f1 := threeValues()
	fmt.Printf("the i1=%d, f1=%f\n", i1, f1)
}

func threeValues() (int, int, float32) {
	return 10, 20, 0.01
}
