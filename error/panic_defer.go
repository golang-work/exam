package main

import (
	"fmt"
)

func fun1(){
    fmt.Println("fun1 text111")
    defer func () {
        fmt.Println("fun1 defer")
    }()
    fun2()
    fmt.Println("fun1 text222")
}

func fun2(){
    fmt.Println("fun2 text111")
    defer func (){
        fmt.Println("fun2 defer")
    }()
    fmt.Println("fun2 text222")
    panic("now panic error")
    fmt.Println("fun2 text333")
}

func fun3(i int) int{
    return i * 10
}

func main() {

    // 所有的 defer 语句都会 保证执行 并把控制权交还给接收到 panic 的函数调用者
    /***** 输出
    在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，
    所有的 defer 语句都会 保证执行 并把控制权交还给接收到 panic 的函数调用者。
    这样向上冒泡直到最顶层，并执行（每层的） defer，
    在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 panicking。

    fun1 text111
    fun2 text111
    fun2 text222
    fun2 defer
    fun1 defer
    panic: now panic error

    goroutine 1 [running]:
    main.fun2()
        /goproject/exam/error/panic_defer.go:24 +0x10f
    main.fun1(0x0)
        /goproject/exam/error/panic_defer.go:13 +0xac
    main.main()
        /goproject/exam/error/panic_defer.go:33 +0x22
    exit status 2
    *****/
    fun1()
}
