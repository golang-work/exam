package main

import "fmt"

const (
	WIDTH = 5
	HEIGHT = 5
)

var screen [WIDTH][HEIGHT]int

func main() {
	fmt.Println(screen)
	for rowIndex, rowColumns := range screen {
		for columnIndex := range rowColumns {
			screen[rowIndex][columnIndex] = rowIndex
		}
	}
	fmt.Println(screen)
}