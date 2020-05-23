package main

import (
	"fmt"
	"math"
)

type Square5 struct {
	side float32
}

type Circle5 struct {
	radius float32
}

type Shaper5 interface {
	Area() float32
}

func main() {
	// 将变量赋值到接口变量或转换成接口变量时
	// 该变量必须要实现该接口
	// 注意实现的接收类型是值类型还是指针类型
	// 如果接收的是指针类型，则在转换时传递地址（&），类型断言时使用指针（*）

	//var areaIntf Shaper5
	//sq1 := &Square5{5}
	//areaIntf = sq1
	//if t, ok := areaIntf.(*Square5); ok {
	//	fmt.Printf("the type of areaIntf is: %T\n", t)
	//}
	//if u, ok := areaIntf.(*Circle5); ok {
	//	fmt.Printf("the type of areaIntf is: %T\n", u)
	//}else{
	//	fmt.Println("areaIntf does not cantain a variable of type circle")
	//}

	// type-switch
	areaIntf := Shaper5(&Square5{5})
	switch t := areaIntf.(type) {
	case *Square5:
		fmt.Printf("the type of areaIntf is: %T\n", t)
	case *Circle5:
		fmt.Printf("the type of areaIntf is: %T\n", t)
	default:
		fmt.Println("areaIntf does not cantain a variable of type circle")
	}
}

func (sq *Square5) Area() float32  {
	return sq.side * sq.side
}

func (ci *Circle5) Area() float32  {
	return ci.radius * ci.radius * math.Pi
}