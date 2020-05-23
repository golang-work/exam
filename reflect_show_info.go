package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo3 struct {
	A int `tag1:"first tag" tag3:"second tag"`
	B string
}

func main() {
	s1 := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo3{
		A : 10,
		B : "salutations",
	}
	fp := &f

	s1Type := reflect.TypeOf(s1)
	gType := reflect.TypeOf(greeting)
	grpType := reflect.TypeOf(greetingPtr)
	fType := reflect.TypeOf(f)
	fpType := reflect.TypeOf(fp)

	examiner(s1Type, 0)
	examiner(gType, 0)
	examiner(grpType, 0)
	examiner(fType, 0)
	examiner(fpType, 0)
}

func examiner(t reflect.Type, depth int)  {
	fmt.Println(strings.Repeat("\t", depth), "type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth + 1), "contained type:")
		examiner(t.Elem(), depth + 1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth + 1), "field", i + 1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth + 2), "tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth + 2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}