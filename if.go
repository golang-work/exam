package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println(runtime.GOOS)
	// 写法一，只有条件表达式：
	//price := 300
	//if price >= 100 {
	//	fmt.Printf("tai gui le\n")
	//}else{
	//	fmt.Printf("nen jie shou\n")
	//}

	// 写法二，初始值 与 条件表达式一起：
	if price := 200; price == 100 {
		fmt.Printf("tai gui le\n")
	} else if price == 200 {
		fmt.Printf("zheng de tai gui le\n")
	} else {
		fmt.Printf("nen jie shou\n")
	}

	// 多返回值函数
	//name := "1212121"
	name := "qing"
	fmt.Printf("%d\n", strconv.IntSize)
	i, err := strconv.Atoi(name)
	if err != nil {
		fmt.Printf("name is not number\n")
		os.Exit(1)
	}

	fmt.Printf("name is:%d\n", i)
}