package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "data/input.dat"
	outputFile := "data/output.dat"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
