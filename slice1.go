package main

import "fmt"

func main(){
	sli := []int{1,2,3,4,5}
	sli[1] = 100

	fmt.Printf("%T, %#v, %v\n", sli, sli, sli)
	fmt.Println(sli, sli[1])
	fmt.Println(&sli, &sli[1])

	fmt.Println()

	var sli2 []int
	fmt.Printf("%T, %v\n", sli2, sli2)

	fmt.Println()

	vara := new(int)
	*vara = 20
	varb := 10

	fmt.Printf("%T, %v, %v\n", vara, vara, *vara)
	fmt.Printf("%T, %v\n", varb, varb)

	fmt.Println("===========================")

	sli3 := []int{1, 3, 5, 6}
	for index, val := range sli3 {
		tmp := 0
		fmt.Printf("temp p = %p\n", &tmp)
		fmt.Printf("val p = %p, val val = %v, val type= %T\n", &val, val, val)
		fmt.Printf("index p = %p, index val = %v, index type= %T\n", &index, index, index)
		fmt.Println()
	}
}
