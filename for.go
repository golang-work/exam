package main

import "fmt"

func main() {
	// 写法一
	//for i :=0; i <10; i++ {
	//	fmt.Printf("i value:%d\n", i)
	//}

	// 写法二
	//i := 0
	//for ; i <10; i++ {
	//	fmt.Printf("i value:%d\n", i)
	//}

	// 写法三
	//for i := 0; ; i++ {
	//	if i == 10 {
	//		break
	//	}
	//	fmt.Printf("i value:%d\n", i)
	//}

	// 写法四
	//i := 0
	//for ; ;  {
	//	if i == 10 {
	//		break
	//	}
	//	fmt.Printf("i value:%d\n", i)
	//	i++
	//}

	// 写法五
	//i := 0
	//for {
	//	if i == 10 {
	//		break
	//	}
	//	fmt.Printf("i value:%d\n", i)
	//	i++
	//}

	// 写法六：基于条件判断的迭代
	i := 0
	for i < 5 {
		fmt.Printf("i value:%d\n", i)
		i++
	}
}
