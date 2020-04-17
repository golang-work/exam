package main

import (
	"fmt"
)

func main() {
	// 写法一，只有条件表达式：
	//dayOfSet := 5
	//
	//switch dayOfSet {
	//case 1:
	//	fmt.Printf("one\n")
	//case 2:
	//	fmt.Printf("two\n")
	//default:
	//	fmt.Printf("other\n")
	//}

	// 写法二，初始值 与 条件表达式一起：
	//switch dayOfSet := 5; dayOfSet {
	//case 1:
	//	fmt.Printf("one\n")
	//case 2:
	//	fmt.Printf("two\n")
	//case 5:
	//	fmt.Printf("in come...\n")
	//default:
	//	fmt.Printf("other\n")
	//}

	// 写法三，case 可以匹配多个值：
	//switch dayOfSet := 5; dayOfSet {
	//case 1:
	//	fmt.Printf("one\n")
	//case 2:
	//	fmt.Printf("two\n")
	//case 3, 4, 5:
	//	fmt.Printf("in come...\n")
	//default:
	//	fmt.Printf("other\n")
	//}

	// 写法四，switch 不提供任何被判断的值（实际上默认为判断是否为 true）
	//dayOfSet := 5
	////switch true{
	//switch{
	//case dayOfSet == 1:
	//	fmt.Printf("one\n")
	//case dayOfSet == 2:
	//	fmt.Printf("two\n")
	//case dayOfSet >= 3 && dayOfSet <=5:
	//	fmt.Printf("in come...\n")
	//default:
	//	fmt.Printf("other\n")
	//}

	// 写法五，switch 只有初始值（注意这个分号不能少）没有条件表达式
	switch dayOfSet := 3; {
	case dayOfSet == 1:
		fmt.Printf("one\n")
	case dayOfSet == 2:
		fmt.Printf("two\n")
	case dayOfSet >= 3 && dayOfSet <=5:
		fmt.Printf("in come...\n")
	default:
		fmt.Printf("other\n")
	}
}