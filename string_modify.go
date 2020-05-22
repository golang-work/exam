package main

import "fmt"

func main() {
	// 如何修改一个字符串
	// string是只读的byte slice（和一些额外的属性）
	str := "this is my go"
	// str[0] = 'T' 这样直接值会报错

	// 需要先转成byte slice，并在需要时把它转换为string类型
	str1 := []byte(str)
	str1[0] = 'T'

	fmt.Println(string(str1))

	// 注意：上面如果对多字节的字符修改会出问题，比如中文 是3个字节
	// 可以将其转成 rune slice
	str2 := "好好学习"
	//str2S := []byte(str2)
	//str2S[0] = '不' // 错误做法，编译不通过，提示大小超出范围
	//fmt.Println(string(str2S))
	str2S := []rune(str2)
	str2S[0] = '不'	// 正确
	fmt.Println(string(str2S))
}