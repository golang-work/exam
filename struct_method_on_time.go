package main

import (
	"fmt"
	"time"
)

// 如果想在非本包的类型上定义方法
// 可以通过别名的方式间接实现
// 因为 类型在其它的或非本地的包时，定义会出错
type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println("full time now:", m.String())
	fmt.Println("first 3 chars:", m.first3Chars())
}