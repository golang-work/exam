package main

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.IsZero())
	fmt.Println(t.Unix())
	fmt.Println(time.Unix(t.Unix(), 0).Format("2006-01-02 15:04:05"))
}
