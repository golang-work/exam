package main

import (
	"fmt"
	"sort"
)

type StuScore struct {
	name string
	score int
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func main()  {
	stus := StuScores{
		{"alan", 95},
		{"hikernell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	fmt.Println("default:\n\t", stus)
	sort.Sort(stus)
	fmt.Println("is sorted?\n\t", sort.IsSorted(stus))
	fmt.Println("sorted:\n\t", stus)
	sort.Sort(sort.Reverse(stus))	// 降序排序，不用修改 less方法里的代码
	fmt.Println("sorted:\n\t", stus)
}
