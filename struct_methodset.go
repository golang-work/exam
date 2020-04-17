package main

import "fmt"

type List []int

func (l List) Len() int {
	return  len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func main() {
	// 值类型
	var lst List
	lst.Append(1)
	fmt.Printf("len:%d\n", lst.Len())

	// 指针类型
	p1st := new(List)
	p1st.Append(2)
	fmt.Printf("len:%d\n", p1st.Len())
}