package main

import "fmt"

type scoreFun func(score int) int

func main() {
	// 定义map:一
	//var mapList map[string]int

	// 定义map:二
	//mapList := map[string]int{"yuwen": 10, "shuxue":20, "yingyu":30}

	// 定义map:三，map值类型可以是任何类型，以下以函数为例
	//mapList := map[string]scoreFun{"yuwen": func(score int) int {
	//	return  score * 1
	//}, "shuxue":func(score int) int {
	//	return  score * 2
	//}, "yingyu":func(score int) int {
	//	return  score * 3
	//}}
	//
	//score := 0
	//for _, value := range mapList {
	//	score++
	//	fmt.Println(value(score))
	//}

	// 定义map:四
	//mapList := make(map[string]int)

	//fmt.Printf("type:%T\n, value:%v\n, goorgvalue:%#v\n", mapList, mapList, mapList)

	// map是引用类型，需要使用make来分配内存或在定义时给定初始值
	mapList1 := map[string]int{"yuwen": 10, "shuxue":20, "yingyu":30}
	fmt.Printf("mapList1 type:%T, mapList1 value:%#v\n",
		mapList1, mapList1)
	mapList2 := mapList1
	mapList2["yuwen"] = 100
	fmt.Printf("mapList2 type:%T, mapList2 value:%#v\n",
		mapList2, mapList2)
	fmt.Printf("mapList1 type:%T, mapList1 value:%#v\n",
		mapList1, mapList1)

	// 在没有分配内存的情况下值是nil
	//var mapList3 map[string]int  // 只定义，未分配内存
	mapList3 := map[string]int{}	// 定义的同时也分配了内存，等同于 mapList3 := make(map[string]int)
	fmt.Printf("mapList3 type:%T, mapList3 value:%#v\n",
		mapList3, mapList3)

}