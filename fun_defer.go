package main

import "fmt"

func main() {
	//function1()

	// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	f()
}

func f()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func function1()  {
	fmt.Printf("in function1 at the top\n")
	defer function2()
	fmt.Printf("in function1 at the bootm!\n")
}

func function2()  {
	fmt.Printf("in function2\n")
}