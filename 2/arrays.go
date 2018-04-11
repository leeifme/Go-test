package main

import (
	"fmt"
	"strconv"
)

func main() {

	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	var grid [2][3]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for _, v := range arr1 {
		fmt.Println(v)
	}
	a := ""
	b, _ := strconv.Atoi(a)
	maxi := -1
	maxv := -1
	for i, v := range arr3 {
		if v > maxv {
			maxi, maxv = i, v
		}
	}
	fmt.Println(maxi, maxv)
	fmt.Println(b)

}
