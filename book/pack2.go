package book

import "fmt"

// 一个包可能有多个 init 函数，它们甚至可以存在于同一个源码文件中
// 它们的执行是无序的
func init()  {
	fmt.Println("book/pack2.go init one...")
}

func init()  {
	fmt.Println("book/pack2.go init two...")
}
