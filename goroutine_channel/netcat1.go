package main

import (
    "io"
    "net"
    "log"
    "os"
    "fmt"
)

func main(){
    conn, err := net.Dial("tcp", "localhost:8000")
    fmt.Println(conn)
    if err != nil {
        log.Fatal(err)
    }
    
    defer conn.Close()
    mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader){
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
