package main

import (
	"fmt"
	"sort"
)

func main()  {
	s := []int{1, 23, 12, 35, 7}
	// 升序
	sort.Ints(s)
	fmt.Println(s)
	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
