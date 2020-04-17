package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "qing "
	fmt.Printf("%s\n", strings.TrimSpace(strings.Repeat(name, 10)))
}