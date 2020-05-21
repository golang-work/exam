package main

import (
	"fmt"
	"syscall"
	"time"
)

func main()  {
	tz, ok := syscall.Getenv("TZ")
	switch{
	case !ok:
		z := time.Local
		fmt.Println(z)
	case tz != "" && tz != "UTC":
		fmt.Println(tz)
	}
}
