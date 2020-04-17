package main

import "fmt"

func main() {
	arrayAges := [...]int{1, 2, 3}
	fmt.Printf("%T, %v\n", arrayAges, arrayAges)
	update(arrayAges)
	fmt.Printf("%T, %v\n", arrayAges, arrayAges)

	for i := 0; i < 5; i++ {
		printArr(&[...]int{i, i * i, i * i * i})
	}
}

func printArr(arr *[3]int) {
	fmt.Println(arr)
}

func update(target [3]int)  {
	target[0] = 100
}