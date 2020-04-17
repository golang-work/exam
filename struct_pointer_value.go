package main

import "fmt"

type B struct {
	think int
}

func (b *B) change() {
	b.think = 1
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	var b1 B
	fmt.Printf("%T, %v, %#v\n", b1, b1, b1)
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B)
	fmt.Printf("%T, %v, %#v\n", b2, b2, b2)
	b2.change()
	fmt.Println(b2.write())
}