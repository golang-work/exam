package main

import "fmt"

func main() {
	sumPrice, err := sumPrice(1, 2, 3)
    if(err == nil){
        fmt.Println(sumPrice)
    }else{
        fmt.Println(err)
    }
}

func sumPrice(i, j, k int) (_ int, err error) {
    cond := false
    if(cond){
        return 0, fmt.Errorf("error %s", "config")
    }
	return  i + j + k, nil
}