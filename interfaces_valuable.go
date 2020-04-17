package main

import "fmt"

type stockPositiong struct {
	ticket string
	sharePrice float32
	count float32
}

func (s stockPositiong) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return  c.price
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable)  {
	fmt.Printf("value of the asset is %f:\n", asset.getValue())
}

func main() {
	var o valueable = stockPositiong{"good", 577.2, 4}
	showValue(o)
	o = car{"bmw", "m3", 66500}
	showValue(o)
}