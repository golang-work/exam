package main

import "fmt"

func init(){
    fmt.Println("init fun...")
    run()
}

func run(){
    fmt.Println("run fun...")
}

var name = func() string {
    fmt.Println("name fun...")
    return ""
}()

func main(){
    fmt.Println("main fun...")
}