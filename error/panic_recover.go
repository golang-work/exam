package main

import (
	"log"
)

// recover 只能在 defer 修饰的函数中使用
// 用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果
// 总结：panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。

func main() {
    protect(func(){
        panic("出错了...")
    })

    log.Println("程序继续执行，do something...")
}

func protect(g func()){
    defer func(){
        log.Println("done")
        if err := recover(); err != nil {
            log.Printf("run time panic:%v", err)
        }
    }()

    log.Println("start")
    g()
}

