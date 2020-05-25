package main

import (
    "fmt"
    "time"
)

func main(){
    /****
    in main()
    about to sleep in main()
    beginning logwait()
    beginning shotWait()
    end of shotWait()
    end of longWait()
    at the end of main()
    ***/

    fmt.Println("in main()")
    go longWait()
    go shotWait()
    fmt.Println("about to sleep in main()")
    time.Sleep(10 * 1e9)
    fmt.Println("at the end of main()")
}

func longWait(){
    fmt.Println("beginning logwait()")
    time.Sleep(5 * 1e9)
    fmt.Println("end of longWait()")
}

func shotWait(){
    fmt.Println("beginning shotWait()")
    time.Sleep(2 * 1e9)
    fmt.Println("end of shotWait()")
}