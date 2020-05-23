package main

import (
	"fmt"
)

type display interface {
	show(s string)
}

type huawei struct {
	text string
}

type sanxin struct {
	text string
}

func (this *huawei) show(s string)  {
	fmt.Println(s)
}

func (this *sanxin) show(s string)  {
	fmt.Println(s)
}

func main() {
	var disI display
	sanx := &huawei{"haha"}
	disI = sanx
	v, ok := disI.(*sanxin);
	fmt.Printf("%#v", v)
	fmt.Println(v, ok)
}