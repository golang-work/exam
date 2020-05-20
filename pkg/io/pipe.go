package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	time.Sleep(30 * time.Second)
}

func PipeWrite(writer *io.PipeWriter){
	data := []byte("i like golang")
	for i := 0; i < 3; i++{
		n, err := writer.Write(data)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("input byte %d\n",n)
	}
	writer.CloseWithError(errors.New("write close"))
}

func PipeRead(reader *io.PipeReader){
	buf := make([]byte, 128)
	for{
		fmt.Println("receive block 5 second...")
		time.Sleep(5 * time.Second)
		fmt.Println("receive int...")
		n, err := reader.Read(buf)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("receive byte: %d\n buf content: %s\n",n,buf)
	}
}
