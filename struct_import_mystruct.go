package main

import (
	"./packme"
	"fmt"
)

func main() {
	s1 := new(packme.ExpStruct)
	s1.Mi1 = 10
	s1.Mf1 = 16.22

	fmt.Printf("Mi1 = %d\n", s1.Mi1)
	fmt.Printf("Mf1 = %f\n", s1.Mf1)
}
