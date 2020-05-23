package main

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (t Celsius) String() string {
	return strconv.FormatFloat(float64(t), 'g', -1, 64) + "°C"
}

func main() {
	data := Celsius(30.3)
	fmt.Println(data)
}