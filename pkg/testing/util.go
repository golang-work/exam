package main

import "fmt"

func MakeStr(prefix string) string{
    return fmt.Sprintf("%s###", prefix)
}