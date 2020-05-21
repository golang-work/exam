package main

import (
	"fmt"
	"strconv"
)

func main(){
	str := strconv.FormatInt(10086, 10)
	fmt.Printf("%T, %s\n", str, str + "haha")
	//fmt.Println(strconv.ParseInt("10", 10, 0))
}
