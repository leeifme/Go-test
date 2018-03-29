package main

import (
	"fmt"
)

func main() {
	n := 0
	sum := 0
	// s := []int{}
	m, _ := fmt.Scan(&n)
	// fmt.Println(n)
	if m == 0 {
		panic("input error!")
	} else {
		s := make([]int, n)
		for i := 0; i <= n; i++ {
			x := 0
			fmt.Scanln(&x)
			s = append(s, x)

		}
		for _, v := range s[5:] {
			if v > 0 {
				sum += v
				if sum == n {
					fmt.Print(1)
				} else if sum > n {
					sum -= v
				}
			}
		}
		fmt.Print(0)
	}
}
