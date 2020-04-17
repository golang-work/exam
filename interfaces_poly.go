package main

import "fmt"

type Shaper1 interface {
	Area() float32
}

type Square1 struct {
	side float32
}

func (sq *Square1) Area() float32 {
	return sq.side * sq.side
}

type Rectangle1 struct {
	length, width float32
}

func (r Rectangle1) Area() float32 {
	return  r.length * r.width
}

func main() {
	r := Rectangle1{5, 3}
	q := &Square1{5}

	//shapes := []Shaper1{Shaper1(r), Shaper1(q)}
	// 或
	shapes := []Shaper1{r, q}
	for n := range shapes {
		fmt.Println("shape deftails:", shapes[n])
		fmt.Println("area of this shape is:", shapes[n].Area())
	}
}