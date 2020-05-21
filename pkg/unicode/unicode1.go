package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(unicode.IsControl('你'))
	fmt.Println(unicode.IsPrint(0))
	fmt.Println(unicode.IsDigit('Ⅷ'))
	fmt.Println(unicode.Is(unicode.Han, '我'))
	fmt.Println(utf8.RuneLen('他'))

}
