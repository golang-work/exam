package main

import "fmt"

var count int

func init()  {
	fmt.Printf("init count is:%d\n", count)
	addCount(20)
	fmt.Printf("init count is:%d\n", count)
}

func addCount(value int)  {
	count += value
}

func main() {
	fmt.Printf("main count is:%d\n", count)
	println()
}