package main

import "fmt"

var i = 5
var str = "ABC"

type Person5 struct {
	name string
	age int
}

type Any interface {

}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value:%v\n", val)
	val = str
	fmt.Printf("val has the value:%v\n", val)
	pers1 := new(Person5)
	pers1.name = "pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value:%v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("type int %T\n", t)
	case string:
		fmt.Printf("type string %T\n", t)
	case bool:
		fmt.Printf("type boolean %T\n", t)
	case *Person5:
		fmt.Printf("type pointer to Person5 %T\n", t)
	default:
		fmt.Printf("unexpected %T\n", t)
	}
}