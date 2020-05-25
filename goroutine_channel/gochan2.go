package main

import (
    "fmt"
    "time"
)

func main(){
    fmt.Println("start...")
    ch := make(chan string)

    go func (){
         time.Sleep(2 * 1e9)
         ch <- "ooooo"
    }()

    fmt.Println(<- ch)
}