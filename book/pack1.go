package book

import "fmt"

var MyInt = 42
var MyFloat = 12.33

var strFun  = func() string {
	fmt.Println("in book/pack1.go define strFun variable ...")
	return ""
}()


func init()  {
	fmt.Println("book/pack1.go init...")
}

func MyFun() string {
	return "Hello my fun"
}