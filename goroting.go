package main

import(
    "fmt"
    "time"
)

func printStr(){    
    for i := 0; i < 10; i++{
        fmt.Println("printStr ... ###")
        time.Sleep(time.Second)
    }
}

func main() {
        go printStr()
        go func(){
            for i := 0; i < 10; i++{
                fmt.Println("main ... ###")
                time.Sleep(time.Second)
            }
        }()
        
        // 程序会一闪而过，因为没有等待机制（等待子线程任务完成）
        fmt.Println("main prosess end")
}