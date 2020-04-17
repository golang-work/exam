package main

import "fmt"

type myInt int

func (i myInt) addOne() myInt {
	return  i + 1
}

func main() {
	v := myInt(10)
	fmt.Println(v.addOne())
}