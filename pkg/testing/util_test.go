package main

import (
	"fmt"
	"testing"
)
func TestMakeStr(t *testing.T){
	tests := []struct{
		input string
		result string
	}{
		{
			"qing1",
			"qing1###",
		},
		{
			"qing2",
			"qing###",
		},
		{
			"qing3",
			"qing3###",
		},
	}

	for _, test := range tests {
		fmt.Println("start...")
		res := MakeStr(test.input)
		if res != test.result {
			t.Errorf("input：%s, result：%s  ->  error", test.input, test.result)
		}
	}
}