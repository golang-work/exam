package main

import (
	"fmt"
	"os"
)

func init(){
    var user = os.Getenv("USER")
    fmt.Println(user)
}

func main() {
    fmt.Println("starting the program")
    panic("a server error occurred:stopping the program!")
    fmt.Println("end the program")
}
