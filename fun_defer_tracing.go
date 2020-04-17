package main

import "fmt"

func main() {
	f1()
}

func trace(s string){
	fmt.Println("entering", s)
}

func untrace(s string){
	fmt.Println("leaving", s)
}

func f1(){
	trace("f1")
	defer untrace("f1")
	fmt.Println("in f1")
	f2()
}

func f2(){
	trace("f2")
	defer untrace("f2")
	fmt.Println("in f2")
}