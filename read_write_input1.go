package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input1 string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input1, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("the input was:%s\n", input1)
	}
}
