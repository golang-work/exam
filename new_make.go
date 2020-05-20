package main

import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// 正常
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// 编译错误：cannot make type Bar，应当使用new
	//z := make(Bar)
	//z.thingOne = "hello"
	//z.thingTwo = 1

	// 正常
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// 运行时错误!! panic: assignment to entry in nil map
	// 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存
	u := new(Foo)
	fmt.Printf("%T  %v\n", u, u)
	//(*u)["x"] = "goodbye"
	//(*u)["y"] = "world"
}
