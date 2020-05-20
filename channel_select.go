package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        c1<- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        c2<- "two"
    }()

    for i := 0; i < 3; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
    // 会一闪而过，没拿到任何chan的数据（因为chan写入端有延迟）
    //    default:
    //        fmt.Println("no data")
        }
    }
}
