package main

import (
	"fmt"
	"sync"
)

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

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)

	// WaitGroup 用于控制多个线程、Goroutine 之间的同步
	var wg sync.WaitGroup
	// 增加内部计数器
	wg.Add(len(chs))

	collect := func(in <-chan int) {
		// Add(-1)——计数器减少1
		defer wg.Done()
		for num := range in {
			out <- num
		}
	}

	// FAN - IN
	for _, ch := range chs {
		go collect(ch)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		// 任何 Goroutine 在调用 Wait 函数时，阻塞等待，直到内部计数器变为 0 时才返回
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// 生产者
	in := producer(1, 2, 3, 4)

	// FAN - OUT 任务分发
	// 处理者
	deal1 := square(in)
	deal2 := square(in)
	deal3 := square(in)

	// FAN -IN 收集处理的数据
	out := merge(deal1, deal2, deal3)

	// 消费者
	for result := range out {
		fmt.Printf("%d ", result)
	}
}
