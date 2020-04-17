package main

import (
	"fmt"
	"math"
)

type Square3 struct {
	side float32
}

type Circle3 struct {
	radius float32
}

type Shaper3 interface {
	Area() float32
}

func main() {
	var areaIntf Shaper3
	sq1 := new(Square3)
	sq1.side = 5

	areaIntf = sq1

	// if 版本
	if t, ok := areaIntf.(*Square3); ok {
		fmt.Printf("the type of areaIntf is: %T\n", t)
	}

	if u, ok := areaIntf.(*Circle3); ok {
		fmt.Printf("the type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	// switch版本
	switch t := areaIntf.(type) {
	case *Square3:
		fmt.Printf("switch the type of areaIntf is: %T\n", t)
	case *Circle3:
		fmt.Printf("switch the type of areaIntf is: %T\n", t)
	default:
		fmt.Println("switch areaIntf does not contain a variable of type Circle")
	}
}

func (sq *Square3) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle3) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}