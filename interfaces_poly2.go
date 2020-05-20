package main

import "fmt"

type Cpu interface{
    Calcu()
}

type Display interface{
    Display()
}

type Memory interface{
    Memory()
}

type Inter struct{
    
}

func (this Inter) Calcu() {
    fmt.Println("inter cpu calcu runing...")
}

type Sansum struct{
    
}

func (this *Sansum) Display() {
    fmt.Println("sansum display runing...")
}

type Aoc struct{
    
}

func (this Aoc) Display() {
    fmt.Println("aoc display runing...")
}

type Kinsdu struct{
    
}

func (this Kinsdu) Memory() {
    fmt.Println("kinsdu memory runing...")
}

type Computer struct{
    cpu Cpu
    display Display
    memory Memory
}

func NewComputer(cpu Cpu, display Display, memory Memory) *Computer{
    return &Computer{cpu:cpu, display:display, memory:memory}
}

func (this Computer) run(){
    this.cpu.Calcu()
    this.display.Display()
    this.memory.Memory()
}

func main() {
	cp1 := NewComputer(Inter{}, &Sansum{}, Kinsdu{})
    cp1.run()
}