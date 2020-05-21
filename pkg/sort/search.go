package main

import (
	"fmt"
	"sort"
)

func main()  {
	x := 11
	s := []int{3, 6, 9, 11, 45}
	pos := sort.Search(len(s), func(i int) bool{
		fmt.Println(i)
		return s[i] >= x
	})

	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置：", pos)
	}else{
		fmt.Println("s 不包含元素", x)
	}
}
