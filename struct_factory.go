package main

import "fmt"

type file struct {
	fd int
	name string
}

func main() {
	f := NewFile(10, "./text.txt")
	fmt.Println(f)
}

func NewFile(fd int, name string) *file {
	if fd < 0 {
		return  nil

	}
	return &file{fd, name}
}