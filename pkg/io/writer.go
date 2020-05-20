package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := os.OpenFile("../text.txt", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件出错")
	}

	//b := []byte("这个世界真美好")
	//file.Write([]byte("abc"))
	//int, err := file.WriteString("###")
	//if err != nil {
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(int)
	//}

	//fmt.Fprintln(file, "哈1")
	fmt.Fprintf(file, "{%s:%v, %s, %v}\n", "price", 10, "count", time.Now().Format("2006-01-02 15:04:05"))

	s := strings.NewReader("haha\n")
	n, err := s.WriteTo(file)
	fmt.Println(n, err)
}
