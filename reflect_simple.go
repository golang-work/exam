package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex byte
	Age int8
}

func main() {
	stu := Student{
		Name:"golang",
		Sex:1,
		Age:56,
	}

	// reflect.TypeOf 返回  reflect.rtype类型的变量
	varType := reflect.TypeOf(stu)

	fmt.Printf("varType 类型: %T, 值：%#v\n", varType, varType)

	// 当 stu 是值类型时，varType.Name() 返回变量类型的名称，
	fmt.Println(varType.Name()) // 返回 Student

	// 当 stu 是指针类型时，varType.Elem() 找出指向或包含的值的类型，值类型时调用varType.Elem()会报错
	// fmt.Println(varType.Elem()) // 返回 main.Student

	// 某些类型（例如切片或指针）没有名称，此方法会返回空字符串
	sli1 := []int{1, 2, 3, 4}
	varType1 := reflect.TypeOf(sli1)
	fmt.Println(varType1.Name())	// 空字符串

	map1 := map[string]string{
		"yuwen":"haha",
		"shuxue":"haoaho",
	}
	varType1 = reflect.TypeOf(map1)
	fmt.Println(varType1.Name())

	fmt.Println("====================================")

	// -Kind 是切片，映射，指针，结构，接口，字符串，数组，函数，int 或其他某种原始类型的抽象表示
	// 类型大的类别名称
	fmt.Println(varType.Kind())	// struct
}