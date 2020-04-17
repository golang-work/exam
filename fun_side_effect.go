package main

import "fmt"

func main() {
	n := 0
	reply := &n
	multiply(10, 6, reply)
	fmt.Println("multiply:", *reply)
	fmt.Println("multiply:", n)
}

func multiply(a, b int, reply *int) {
	*reply = a * b
}
