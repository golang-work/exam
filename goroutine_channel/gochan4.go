package main

import (
	"fmt"
	"time"
	_ "runtime"
)

func main() {
    //runtime.GOMAXPROCS(8) // 最多同时使用2个核

	ch1 := make(chan string)

	for i := 1; i <= 10; i++ {
		go addData33(ch1, i)
	}

	for k := 0; k < 50; k++ {
		fmt.Println(<- ch1)
	}
}

func addData33(ch chan<- string, i int) {
	for j := 1; j <= 5; j++ {
		time.Sleep(1 * 1e9)
		ch <- fmt.Sprintf("线程%d：%d", i, j)
	}
}