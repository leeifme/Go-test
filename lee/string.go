package main

import (
	"fmt"
)

type SortInterface interface {
	sort()
}
type Sortor struct {
	name string
}

func main() {
	var arry []int
	fmt.Println("Please input your full name: ")
	fmt.Scanln(&arry)
	learnsort := Sortor{name: "快速排序--从小到大--不稳定--nlog2n最坏n＊n---"}
	learnsort.sort(arry)
	fmt.Println(learnsort.name, arry)
}

func (sorter Sortor) sort(arry []int) {
	if len(arry) <= 1 {
		return //递归终止条件，slice变为0为止。
	}
	mid := arry[0]
	i := 1 //arry[0]为中间值mid，所以要从1开始比较
	head, tail := 0, len(arry)-1
	for head < tail {
		if arry[i] > mid {
			arry[i], arry[tail] = arry[tail], arry[i] //go中快速交换变量值，不需要中间变量temp
			tail--
		} else {
			arry[i], arry[head] = arry[head], arry[i]
			head++
			i++
		}
	}
	arry[head] = mid
	sorter.sort(arry[:head]) //这里的head就是中间值。左边是比它小的，右边是比它大的，开始递归。
	sorter.sort(arry[head+1:])

}
