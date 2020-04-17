package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	printValues()
	numx2, numx3 = getX2AndX3_2(num)
	printValues()
	numx2, numx3 = getX2AndX3_3(num)
	printValues()
}

func printValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num= %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// 其实在定义返回值及类型时，变量名已经有了定义及对应类型的零值
	//return x2, x3
	return
}

func getX2AndX3_3(input int) (x2 int, x3 int) {
	return 2 * input, 3 * input
}
