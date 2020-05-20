package main

import(
    "fmt"
)

func main() {
    cha1 := make(chan string, 10)
    cha1<- "hello"
    cha1<- "world"
    close(cha1)
    //cha1<- "other"
    
    fmt.Printf("%v\n", cha1)
    fmt.Printf("len:%d, cap:%d\n", len(cha1), cap(cha1))
    
    //v, ok := <-cha1
    //for i := 0; i < 100; i++{
    //    v, ok = <-cha1
    //}    
    //fmt.Printf("value = %v, ok state = %t\n", v, ok)
    
    for v := range cha1{
        fmt.Printf("value = %v\n", v)
    }
}