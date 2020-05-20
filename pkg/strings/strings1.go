package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "this is str"
	strs := strings.Split(str, " ")
	fmt.Println(strs)

	str2 := []string{"p p a", "j j b"}
	fmt.Println(strings.Join(str2, "+"))
	fmt.Println(str2[1])

}