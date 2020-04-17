package main

import "fmt"

type Simpler interface {
	Get() int
	Set(i int)
}

type Product struct {
	price int
}

func (p *Product) Set(i int) {
	p.price = i
}

func (p Product) Get() int {
	return p.price
}

func simpler(s Simpler)  {
	s.Set(10)
	fmt.Println(s.Get())
}

func main() {
	s := Simpler(&Product{})
	simpler(s)
}