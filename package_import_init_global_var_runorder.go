package main

import (
	"./book"
	"fmt"
)

var strFun  = func() string {
	fmt.Println("in main define strFun variable ...")
	return ""
}()

func init()  {
	fmt.Println("main init1...")
}

func init()  {
	fmt.Println("main init2...")
}


func main() {
	fmt.Println(book.MyFun())
}
