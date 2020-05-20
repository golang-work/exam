package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("this 我s my golang...")
	reader.Seek(5, io.SeekStart)

	s := make([]byte, 100)
	reader.Read(s)
	fmt.Printf("%s\n", s)
}
