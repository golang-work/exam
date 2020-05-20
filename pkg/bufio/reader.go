package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main(){
	//in := os.Stdin
	//con := make([]byte, 10)
	//in.Read(con)
	//fmt.Printf("%s\n", con)


	//r := bufio.NewReader(in)

	reader := bufio.NewReader(strings.NewReader("http://www.baidu.com"))
	line, _ := reader.ReadBytes(':')
	fmt.Printf("the line：%s\n", line)

	n, _ := reader.ReadBytes(':')
	fmt.Printf("the line：%s\n", n)

	fmt.Println(string(n))
}
