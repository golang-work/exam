package main

import "fmt"

func main() {
	f := func ()  {
		fmt.Println("aaaa")
	}

	f()
	fmt.Printf("%T, %#v, %p", f, &f, f)

	//func(){
	//	fmt.Println("bbb")
	//}()


}

