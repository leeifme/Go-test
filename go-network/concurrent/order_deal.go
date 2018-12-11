package concurrent

import "fmt"

type OrderInfo struct {
	ID int
}

func producer(order chan<- OrderInfo) {
	for i := 0; i < 7; i++ {
		order <- OrderInfo{ID: i * i}
		fmt.Println("生成订单号：", i*i)
	}
	close(order)
}

func consumer(deal <-chan OrderInfo) {
	// for {
	// 	if num, ok := range con; ok {
	// 		fmt.Println("消费者，消费：", num)
	// 	} else {
	// 		fmt.Println("消费完毕")
	// 		break
	// 	}
	// }
	for order := range deal {
		fmt.Println("处理订单号：", order.ID)
	}
}

func orderdeal() {
	ch := make(chan OrderInfo, 5)

	go producer(ch)
	consumer(ch)
}
