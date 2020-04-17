package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.firstName)
}

func main() {
	// 值的方式
	var p1 Person
	p1.firstName = "Chris"
	p1.lastName = "Woodward"
	upPerson(&p1)
	fmt.Printf("p1 type:%T, value:%v\n", p1, p1)

	// 指针方式
	p2 := new(Person)
	p2.firstName = "Chris"
	p2.lastName = "Woodward"
	(*p2).lastName = "Woodward"
	upPerson(p2)
	fmt.Printf("p2 type:%T, value:%v\n", p2, p2)

	// 字面意义
	p3 := &Person{"Chris", "Woodward"}
	upPerson(p3)
	fmt.Printf("p3 type:%T, value:%v\n", p3, p3)

	fmt.Println(unsafe.Sizeof(Person{}))
}
