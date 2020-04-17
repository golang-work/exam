package main

import "fmt"

type struct2 struct {
	i1 int
	f1 float64
	str string
}

func main() {
	// 结构体指针类型 在复制到另一个变量时，复制的是地址，一改全改
	//ms1 := &struct2{1, 1.0, "a"}
	//ms2 := ms1
	//ms2.i1 = 100

	// 结构体类型 在复制到另一个变量时，复制的是值的拷贝，改时不影响原来的值
	ms1 := struct2{1, 1.0, "a"}
	ms2 := ms1
	ms2.i1 = 100

	fmt.Printf("ms1 type:%T, value:%v\n", ms1, ms1)
	fmt.Printf("ms2 type:%T, value:%v\n", ms2, ms2)
}
