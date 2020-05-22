package main

import "fmt"

func main() {
	scores := [...]int{10, 20, 30}
	fmt.Println("更新之前", scores)
	updateSc(&scores)
	fmt.Println("更新之后", scores)
}

func updateSc(scores *[3]int)  {
	for i := range scores {
		//scores[i] *= 10 // 自动解引用
		(*scores)[i] *= 10
	}
	fmt.Println("函数体内更新后", scores)
}