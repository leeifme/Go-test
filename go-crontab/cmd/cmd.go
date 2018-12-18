package main

import (
	"fmt"
	"time"
	"os/exec"
	"context"
)

type result struct{
	outPut []byte
	err error
}

func main() {

	var(
		ctx context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
		cmd *exec.Cmd
		res *result
	)
	//执行1个cmd,让它在一个协程里去执行，让它执行2秒,1秒的时候，我们杀死cmd
	//sleep 2;ehco hello world
	

	//context 有一个chan byte
	//cancelFunc:  关闭 close(chan byte)
	ctx , cancelFunc = context.WithCancel(context.TODO())

	//创建一个结果队列
	resultChan = make(chan *result,100)

	go func () {
		var (
			outPut []byte
			err error
		)

		cmd = exec.CommandContext(ctx, "C:\\Program Files\\Software\\Git\\bin\\bash.exe", "-c","sleep 2;ehco hello world")
		outPut ,err = cmd.CombinedOutput()

		resultChan <- &result{
			outPut: outPut,
			err: err,
		}
	}()

	//继续往下走
	time.Sleep(1*time.Second)

	//取消上下文
	cancelFunc()

	//在main协程里，等待子协程的退出，并打印任务执行结果
	res = <- resultChan

	if res.err != nil{
		fmt.Println("err: ",res.err)
	}else {
		fmt.Println("cmd output: ",res.outPut)
	}
}
