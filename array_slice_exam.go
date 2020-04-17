package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("%#v", b[1:4])
}