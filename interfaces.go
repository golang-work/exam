package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter(param string) string
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter(param string) string {
	return param
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	// 接口变量里包含了接收者实例的值和指向对应方法表的指针
	//var areaIntf Shaper
	//areaIntf = sq1

	// 简写 1
	areaIntf := Shaper(sq1)
	fmt.Printf("areaIntf type:%T value:%#v\n", areaIntf, areaIntf)

	// 简写 2
	//areaIntf := sq1
	fmt.Printf("the square has area:%f\n", areaIntf.Area())
	fmt.Printf("the square has perimeter:%s\n", areaIntf.Perimeter("###1"))
}