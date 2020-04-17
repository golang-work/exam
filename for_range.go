package main

import "fmt"

func main() {
	name := "qing"
	for key, value := range name {
		fmt.Printf("%d => %c\n", key, value)
	}
}
