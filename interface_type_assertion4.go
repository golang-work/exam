package main

import "fmt"

type Simpler1 interface {
	Get()int
	Set(int)
}

type Simple struct {
	count int
}

type RSimple struct {
	count int
}

func (this Simple) Get() int {
	return this.count
}

func (this *Simple) Set(c int) {
	this.count = c
}

func (this RSimple) Get() int {
	return this.count
}

func (this *RSimple) Set(c int) {
	this.count = c
}

func exec(sI Simpler1, count int)  {
	sI.Set(count)
	fmt.Println(sI.Get())
}

func fi(si Simpler1)  {
	switch si.(type) {
	case *Simple:
		fmt.Println("si is simple instance")
	case *RSimple:
		fmt.Println("si is rsimple instance")
	}
}

func main() {
	s := Simple{}
	rs := RSimple{}
	exec(&s, 10)
	exec(&rs, 20)

	//fi(Simpler1(&s))
	fi(Simpler1(&rs))
}