package main

import "fmt"

func main() {
	slice1 := []string{"这个世界真", "美好"}
	fmt.Printf("%#v\n", slice1)

	for _, value := range slice1 {
		fmt.Printf("%#v\n", value)
	}
}