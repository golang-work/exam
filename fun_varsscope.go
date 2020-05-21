package main

import "fmt"

var slice []func()

func main() {
    sli := []int{1, 2, 3, 4, 5}

    fmt.Println(sli)

    for _, v := range sli {
        //slice = append(slice, func(){
        //    fmt.Println(v * v) // 直接打印结果
        //})
	app(v)

    }

    //fmt.Println(v)

    fmt.Println(slice)

    for _, val := range slice {
        val()
    }

	for i := 0; i < 5; i++ {
		j := 0
		fmt.Println(&i, &j)
	}
}

func app(v int){
	slice = append(slice, func(){
            fmt.Println(v * v) // 直接打印结果
        })
}
