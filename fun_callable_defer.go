package main

import "fmt"

func main() {
	fmt.Println(f3())
}

func f3() (ret int) {
	defer func(){
		ret++
	}()
	ret = 1
	return
}

