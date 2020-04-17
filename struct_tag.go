package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool "an important answer"
	field2 string "the name of the thing"
	field3 int "how much there are"
}

func main() {
	tt := TagType{true, "barak obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int){
	fmt.Printf("item %d type:%s, name:%s, tag:%s\n",
		ix,
		reflect.TypeOf(tt).Field(ix).Type,
		reflect.TypeOf(tt).Field(ix).Name,
		reflect.TypeOf(tt).Field(ix).Tag)
}
