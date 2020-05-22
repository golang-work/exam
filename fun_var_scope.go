package main

import "fmt"

var total = 10

func Adder(def int) (func(i int) int, int) {
	// 闭包里如果使用了外部变量，会一直保留在返回的闭包里
	sum := def
	return func(i int) int {
		sum += i
		return sum * total
	}, sum
}
func main() {
	fun1, sum := Adder(1)
	fmt.Println(sum)
	for i := 0; i < 5; i++ {
		fmt.Println(fun1(i))
	}
	fmt.Println(sum)

	fmt.Println()

	fun2, sum := Adder(2)
	total = 20	// 此时修改了全局变量里的值时，fun2闭包里也会随着变化
	fmt.Println(sum)
	for i := 0; i < 5; i++ {
		fmt.Println(fun2(i))
	}
	fmt.Println(sum)
}