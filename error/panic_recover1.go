package main

import (
	"fmt"
)

// 总结：panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。
// 如果堆栈内有多个defer recover时，最先出栈的生效
// 触发panic的函数内如果有捕获recover，该函数 触发panic的 后面代码不执行

func badCall(){
    defer func(){
        fmt.Println("###111")
    }()

    defer func(){
        if e := recover(); e != nil {
            fmt.Println("###222")
        }
    }()

    panic("bad end")

    fmt.Println("###333")
}

func test(){
    defer func(){
        if e := recover(); e != nil {
            fmt.Printf("panicing %s, \n", e)
        }
    }()

    badCall()
    fmt.Printf("after bad call\n")
}

func main() {
    fmt.Printf("calling test\n")
    test()
    fmt.Printf("test completed\n")
}
