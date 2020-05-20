package main

import "fmt"

type List1 []int

func (l List1) Len() int {
	return  len(l)
}

func (l *List1) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start , end int)  {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len() * 10 > 42
}

func main() {
	var lst List1

	// 使用接口时：这样直接调用会出错，因为 Append 方法定义在了指针上
	//CountInto(lst, 1, 10)
	// 可以 CountInto(&lst, 1, 10) 解决


	if LongEnough(lst) {
		fmt.Printf("- lst is long enough")
	}

	plst := new(List1)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough")
	}
}