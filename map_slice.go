package main

import "fmt"

func main() {
	items := make([]map[string]int, 3)
	for i := range items {
		items[i] = make(map[string]int, 1)
		items[i]["score"] = 20
	}
	fmt.Printf("version A:%#v\n", items)

	items2 := make([]map[string]int, 5)
	for _, item := range items2 {
		item = make(map[string]int, 1)
		item["score"] = 30
	}
	fmt.Printf("version B:%#v\n", items2)

	// 需要注意的是，应当像 A 版本那样通过索引使用切片的 map 元素。
	// 在 B 版本中获得的项只是 map 值的一个拷贝而已，所以真正的 map 元素没有得到初始化。

	data := []byte("Go语言")
	fmt.Printf("%T %v\n", data, data)
}