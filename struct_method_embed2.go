package main

import "fmt"

type Log1 struct {
	msg string
}

type Customer1 struct {
	Name string
	Log1
}

func main() {
	c := &Customer1{"barak", Log1{"1 - yes we can!"}}
	c.Add("2 - after me the world will be a beatter!")
	fmt.Println(c)
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

func (c Customer1) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log1)
}

func (l *Log1) String() string {
	return  l.msg
}