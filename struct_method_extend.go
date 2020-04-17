package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type NamedPoint struct {
	Point
	name string
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

// 重写继承来的方法
//func (p NamedPoint) Abs() float64 {
//	return p.x * p.x + p.y * p.y
//}

func main() {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}