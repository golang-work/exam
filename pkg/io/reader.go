package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func main() {
	//var err error
	//data, err := ReadFrom(os.Stdin, 11)
	//if err == nil {
	//	fmt.Println(data)
	//	for _, value := range data {
	//		fmt.Println(value)
	//	}
	//}

	//file, err := os.Open("../text.txt")
	//if err != nil {
	//	fmt.Println("打开文件 text.txt 错误")
	//	os.Exit(1)
	//}
	//
	//data, err := ReadFrom(file, 6)
	//if err == nil {
	//	fmt.Printf("%s\n", data)
	//}

	//data, _ := ReadFrom(strings.NewReader("abcdef"), 3)
	//fmt.Println(data)

	//fmt.Println(io.EOF.Error())

	file, err := os.Open("../text.txt")
	if err != nil {
		fmt.Println("打开文件出错")
	}

	b := make([]byte, 10)
	n, err := file.ReadAt(b, 4)
	fmt.Printf("%s\t\t\t%d\n", b, n)
}
