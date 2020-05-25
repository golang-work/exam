package main

import (
	"fmt"
)

// 对于 命名的返回变量 ，在 defer内可以继续修改
func fun1(i int) (i2 int) {
    i2 = i
    defer func () {
        i2 *= 20
        fmt.Println("fun1 defer...", i2)
    }()

    fmt.Println("fun1 return...", i2)
    return
}

func main() {
    fmt.Println(fun1(10))
}
