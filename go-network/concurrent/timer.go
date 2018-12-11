package concurrent

import (
	"fmt"
	"time"
)

func timer() {
	// 普通定时器
	commontimer()

	// 定时器停止和重置
	stopandreset()
}

func stopandreset() {
	fmt.Println("now:   ", time.Now())
	myTimer := time.NewTimer(time.Second * 5)

	go func() {
		<-myTimer.C // 默认读——阻塞，定时时长到达后，系统写入当前时间，解除阻塞
		fmt.Println("定时时间到：", time.Now())
	}()
	//myTimer.Stop()
	myTimer.Reset(time.Second * 1)
	for {

	}
}

func commontimer() {
	fmt.Println("now:   ", time.Now())

	time := <-time.After(time.Second * 2)

	fmt.Println("定时后:", time)
}
