package main

import "fmt"

type Action interface {
	Submit()
}

type Article struct {

}

func (this *Article) Submit() {
	fmt.Println("ok...")
}

func main() {
	art := Article{}
	actI := Action(&art)
	if _, ok := actI.(*Article); ok {
		fmt.Println("aI instance of article")
	}
}