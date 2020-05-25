package main

import (
	"fmt"
)

type AbortError struct {
	Code int
	Msg  string
}

func (e *AbortError) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Msg)
}

func main() {
	abort := &AbortError{
		Code: 100010,
		Msg:  "参数错误",
	}

	err := error(abort)

	//if err, ok := err.(*AbortError); ok {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(err.Error())
	//}

	switch err := err.(type) {
	case *AbortError:
		fmt.Println(err)
	default:
		fmt.Println(err)
	}
}
