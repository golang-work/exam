package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()
	fmt.Printf("a type:%T, value:%v, canadd:%t\n", a, a, a.CanAddr())
	fmt.Printf("b type:%T, value:%v, canadd:%t\n", b, b, b.CanAddr())
	fmt.Printf("c type:%T, value:%v, canadd:%t\n", c, c, c.CanAddr())
	fmt.Printf("d type:%T, value:%v, canadd:%t\n", d, d, d.CanAddr())
}