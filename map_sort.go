package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{
		"alpha":   34,
		"bravo":   56,
		"charlie": 23,
		"delta":   87,
		"echo":    56,
		"foxtrot": 12,
		"golf":    34,
		"hotel":   16,
		"indio":   87,
		"juliet":  65,
		"kilo":    43,
		"lima":    98,
	}
)

func main() {
	fmt.Println("unsorted:")
	for key, value := range barVal {
		fmt.Printf("key:%v, value:%v\n", key, value)
	}

	keys := make([]string, len(barVal))
	i := 0
	for key := range barVal {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, key := range keys {
		fmt.Printf("key:%v, value:%v\n", key, barVal[key])
	}
}
