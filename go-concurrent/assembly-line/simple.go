package main

import "fmt"

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	return out
}

func main() {
	in := producer(1, 2, 3, 4)
	out := square(in)
	for result := range out {
		fmt.Printf("%d ", result)
	}
}
