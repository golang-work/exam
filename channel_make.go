package main

import(
    "fmt"
)

func main() {
        // channel 是指针类型
        //var cha1 chan int
        //cha1<- 10
        
        cha1 := make(chan int, 5)
        cha1<- 10
        fmt.Printf("%T, %v, %p\n", cha1, cha1, &cha1)
        fmt.Printf("len:%d, cap:%d\n", len(cha1), cap(cha1))
        
        num1 := <-cha1
        fmt.Printf("%d\n", num1)
        
        fmt.Printf("len:%d, cap:%d\n", len(cha1), cap(cha1))
}