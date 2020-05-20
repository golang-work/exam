package main

import (
	"bytes"
	"fmt"
)

func main()  {
	//slice1 := []byte("我的")
	//fmt.Println(bytes.Contains(slice1, []byte("我")))

	b := []byte("欢迎学习")
	for i, value := range b {
		fmt.Printf("%d => %v\n", i, string(value))
	}
	fmt.Println()
	r := bytes.Runes(b)
	for i, value := range r {
		fmt.Printf("%d => %v\n", i, string(value))
	}

}
