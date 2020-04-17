package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

func main() {
	// 写法一
	//c := new(Customer)
	//c.Name = "barak"
	//c.log = new(Log)
	//c.log.msg = "1 - yes we can!"

	// 写法二
	c := &Customer{"barak", &Log{"1 - yes we can!"}}
	// fmt.Println(c)
	c.Log().Add("2 - after me the world will be a better place!")
	c.Log().Add("3 - other...!")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (c Customer) Log() *Log {
	return c.log
}

func (l Log) String() string {
	return  l.msg
}