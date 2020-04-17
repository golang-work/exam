package main

import "fmt"

func main() {
	p2 := add2()
	fmt.Printf("call add2 for 2 give:%v\n", p2(2))

	adder := adder(2)
	fmt.Printf("call adder:%v\n", adder(2))
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int{
	return  func(b int) int{
		return  a + b
	}
}