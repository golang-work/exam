package main

import (
    "fmt"
    "time"
    "log"
    "net/http"
)

func main(){
    fmt.Println("please visit http://127.0.0.1:8686")
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
        s := fmt.Sprintf("hello, go! -- Time:%s", time.Now().String())
        fmt.Fprintf(w, "%v\n", s)
        log.Printf("%v\n", s)
    })
    if err := http.ListenAndServe(":8686", nil); err != nil {
        log.Fatal("listenandserver:", err)
    }
}