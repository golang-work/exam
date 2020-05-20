package main

import "fmt"

func main(){
	str := []byte("中文")
	fmt.Println(len(str))
	fmt.Printf("%b\n", 100)
	fmt.Printf("%c\n", 1009823)
	fmt.Printf("%d\n", 1100100)
	fmt.Printf("%o\n", 1100100)
	fmt.Printf("%x\n", 1100100)
	fmt.Printf("%X\n", 1100100)
	fmt.Printf("%s\n", []rune("中文"))
}
