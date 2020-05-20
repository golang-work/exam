package main

import "fmt"

type Task struct{
     f func() error
}

func NewTask(f func() error) *Task{
    return &Task{f:f}
}

type workChan interface{
    
}

type Pool struct{
    workChan chan *Task
    entryChan chan *Task
    countNumber int
}

func NewPool(){

}

func (p *Pool) run(){
    for i := 0; i < p.countNumber; i++{
        go 
    }
}

func main() {
	task := NewTask(func() error{
        fmt.Println("task 1")
        return nil
    })
    _  = task.f()
}
