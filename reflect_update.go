package main

import (
	"fmt"
	"reflect"
)

type Foo6 struct {
	Name string
}

func main() {
	f6 := Foo6{"php"}
	fmt.Printf("f6 update befor: %v\n", f6)

	// 如果希望能够使用反射来修改值，则必须传递地址: refPtrVal := reflect.ValueOf(&f6)
	// 否则只能读取
	refPtrVal := reflect.ValueOf(&f6)
	fmt.Println(refPtrVal.Type(), refPtrVal.Kind())

	// 修改值
	// 1. 使用指针
	// 需要三个步骤：
	// 第一步 是调用Addr()方法，它返回一个Value，里面保存了指向变量的指针
	// 第二步 在Value上调用Interface()方法，也就是返回一个interface{}，里面包含指向变量的指针
	// 第三步 如果我们知道变量的类型，我们可以使用类型的断言机制将得到的interface{}类型的接口强制转为普通的类型指针
	px := refPtrVal.Elem().Addr().Interface().(*Foo6)
	(*px).Name = "golang ptr"
	// 2. 不使用指针
	//refPtrVal.Elem().Field(0).Set(reflect.ValueOf("golang func"))

	fmt.Printf("f6 update after: %v\n", f6)
}