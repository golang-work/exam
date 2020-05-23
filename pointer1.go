package main

import "fmt"

type Book struct {
	title string
}

func main() {
	var str string
	str = "qing"

	strP := &str
	fmt.Printf("%T, %#v", strP, strP)

	fmt.Println()

	bk := &Book{"golang"}
	fmt.Printf("%T, %#v, %p, %p", bk, bk, bk, &bk)
}