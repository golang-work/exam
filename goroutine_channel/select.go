package main

import (
    "fmt"
    "time"
)

func main(){
    c1, c2 := make(chan int, 1), make(chan int, 1)
    go func(){
        time.Sleep(2 * time.Second)
        c1 <- 1
    }()
    go func(){
        time.Sleep(time.Second)
        c2 <- 2
    }()

    //for i := 0; i < 2; i++ {
        select{
        case c1 <- 3:
            fmt.Println("send ch1 data by 3")
        case msg1 := <- c1:
            fmt.Println("received msg from c1", msg1)
        case msg2 := <- c2:
            fmt.Println("received msg from c2", msg2)
        default:
            fmt.Println("received default msg")
        }
    //}

    fmt.Println("done...")
}