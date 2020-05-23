package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type of v:", v.Type())

	// 直接设置会出错
	//v.SetFloat(3.1415)

	// 需要先检查是否可以设置，若不可直接设置通过间接的方式
	fmt.Println("settability of v:", v.CanSet()) // false
	v = reflect.ValueOf(&x)	// 引用传递
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet()) // false
	v = v.Elem()
	fmt.Println("the elem of v is", v)
	fmt.Println("settability of v:", v.CanSet()) // true
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}