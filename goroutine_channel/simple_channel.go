package main

import (
    "fmt"
    "time"
)

func main(){
    // 所有的类型都可以用于通道，空接口 interface{} 也可以。甚至可以（有时非常有用）创建通道的通道
    // channel 是指针类型，所以使用前需要用make分配内存

    // 方式一
    //var ch1 chan string
    //ch1 = make(chan string)

    方式二
    cha1 := make(chan string)
}
