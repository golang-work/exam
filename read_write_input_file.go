package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//var inputFile *os.File
	//var inputError, readerError os.Error
	//var inputReader *bufio.Reader
	//var inputString string

	inputFile, inputError := os.Open("data/input.dat")
	if inputError != nil {
		fmt.Printf("an error occurred on openging the inputifl\n" +
			"does the file exist?\n" +
			"have you got acces to it?\n")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for{
		inputString, readError := inputReader.ReadString('\n')
		fmt.Printf("the input was:%s", inputString)
		if readError == io.EOF {
			return
		}
	}
}
