package main

import (
	"fmt"
	_ "time"
)

type Sender interface {
	Send(t string)
}

type Mail struct {
}

type Sms struct {
}

type Person5 struct {
	name string
	age  int
}

func (mail *Mail) Send(t string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%s mail send...%d\n", t, i)
		for j := 0; j < 9000000000; j++ {

		}
		//time.Sleep(time.Second)
	}
	pc <- true
}

func (sms *Sms) Send(t string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s sms send...%d\n", t, i)
		for j := 0; j < 9000000000; j++ {

		}
		//time.Sleep(time.Second)
	}
	pc <- true
}

func (person *Person5) SendNotice(sender []Sender) {
	for _, s := range sender {
		go s.Send("qa")
		go s.Send("in")
		go s.Send("up")
	}
}

var pc = make(chan bool, 6)

func main() {
	p5 := Person5{name: "qing", age: 20}

	mail := &Mail{}
	sms := &Sms{}
	senderList := []Sender{Sender(mail), Sender(sms)}

	p5.SendNotice(senderList)

	for i := 0; i < 6; i++ {
		<-pc
	}

	fmt.Println("main process...end")
}
