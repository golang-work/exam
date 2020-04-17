package main

import (
	"fmt"
	"regexp"
)

func main() {
	//fmt.Println([]rune("中文汉字"))
	str := "hello qing 15151510866"
	pat := "[0-9]+"
	matched, err := regexp.Match(pat, []byte(str))
	fmt.Printf("%v, %v\n", matched, err)
	if matched {
		re, _ := regexp.Compile(pat)
		str1 := re.ReplaceAllString(str, "mobile")
		fmt.Printf("str1 value:%v\n", str1)

		str2 := re.ReplaceAllStringFunc(str, func(s string) string {
			fmt.Printf("callable s value:%v\n", s)
			return  "####"
		})
		fmt.Printf("str2 value:%v\n", str2)
	}
}
