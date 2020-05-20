package main

import(
    "fmt"
)

type valList interface{}

func main() {
    var (
        //num1 = "abc"
        num1 = 10
        num2 int
    )
    var l1 valList
    l1 = num1
    if v, ok := l1.(int); ok{
        num2 = v
        fmt.Printf("num2 type:%T, value:%v\n", num2, num2)
    }
    
    
    //l1 := valList(10)
    
    // 
    //l1 := []valList{10, "abc", 3.34}
    //for _, v := range l1 {
    //    switch v.(type){
    //    case int:
    //        fmt.Printf("int = %d\n", v)
    //   case string:
    //        fmt.Printf("string = %v\n", v)
    //    default:
    //        fmt.Printf("unknow type...\n")
    //    }
    //}
    
    fmt.Printf("l1 type:%T, value:%v\n", l1, l1)
}