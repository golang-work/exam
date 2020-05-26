package main

import (
	"fmt"
)

func main() {
    for _, s := range `-|/\` {
        fmt.Printf("%c\n", s)
    }
}