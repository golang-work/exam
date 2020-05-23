package main

import "fmt"

type myInt2 int

func main() {
	data := myInt2(10)
	data2 := int(10)

	fmt.Printf("%t", data == myInt2(data2))


}