package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	chcom := make(chan int)

	go addData11(ch1)
	go getData22(ch1, chcom)

	//time.Sleep(3 * 1e9)
	<-chcom
}

func addData11(ch chan<- int) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * 1e9)
		ch <- i
	}
}

func getData22(ch <-chan int, c chan<- int) {
	for {
		fmt.Println(<-ch)
	}

	// 死锁
	c <- 1

}
