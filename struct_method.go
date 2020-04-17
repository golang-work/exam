package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("two1 the sum is:%d\n", two1.AddThem())
	fmt.Printf("two1 add them to the param:%d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("two2 the sum is:%d\n", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

//func (tn TwoInts) AddThem() int {
//	return tn.a + tn.b
//}

func (tn *TwoInts) AddToParam(param int) int {
	return  tn.a + tn.b + param
}
