package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

func (t T) String() string {
	return strconv.Itoa(t.a) + "/" +
		strconv.FormatFloat(float64(t.b), 'g', -1, 32) + "/" +
		t.c
}

func main() {
	fmt.Println(T{10, -2.35, "golang"})
}