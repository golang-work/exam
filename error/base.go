package main

import (
    "fmt"
    "errors"
)

func main(){
    err := errors.New("math - square root of negative number")
    fmt.Printf("err:%v\n", err)
}