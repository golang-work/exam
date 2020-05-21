package main

import (
	"fmt"
	"time"
)

func main()  {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-05-21 09:14:00", time.Local)
	fmt.Println(time.Now().Sub(t).Minutes())

	fmt.Println(t.Truncate(1 * time.Hour))
	fmt.Println(t.Round(1 * time.Minute))
}
