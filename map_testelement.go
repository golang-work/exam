package main

import "fmt"

func main() {
	mapList := map[string]int{
		"yuwen":10,
		"shuxue":20,
		"yingyu":30,
	}
	fmt.Printf("%v\n", mapList)

	// 如果 map 中不存在 other，other 就是一个值类型的空值
	fmt.Printf("%v\n", mapList["other"])

	// 为了区分到底是 other 不存在还是它对应的 other 就是空值
	value, isExists := mapList["other"]
	fmt.Printf("%v , %t\n", value, isExists)

	// 删除key， 如果 key 不存在，该操作不会产生错误
	delete(mapList, "yuwen")
	fmt.Printf("%v\n", mapList)
}