// 存在2个同名方法但不在同一个接口时
// 结论：同名方法的接口都实现了
package main

import "fmt"

type Car interface {
	run()
}

type Bike interface {
	run()
}

type Tool struct {

}

func (t Tool) run()  {
	fmt.Println("run...")
}

func main()  {
	tool := Tool{}
	tool.run()
}