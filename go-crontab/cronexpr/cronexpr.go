package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {

	var (
		nowTime      time.Time
		expr         *cronexpr.Expression
		cronJob      *CronJob
		cronJobTable map[string]*CronJob
	)
	// 创建任务调度列表
	cronJobTable = make(map[string]*CronJob, 0)
	nowTime = time.Now()
	fmt.Println("begin time: ", nowTime)

	expr = cronexpr.MustParse("*/2 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(nowTime),
	}
	cronJobTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(nowTime),
	}
	cronJobTable["job2"] = cronJob

	// 创建一个调度协程
	go func() {
		var (
			cronJobName string
			// cronJob     CronJob
		)
		// 轮询检查 所定时任务 是否该被调度执行
		for {
			nowTime = time.Now()
			for cronJobName, cronJob = range cronJobTable {
				if cronJob.nextTime.Before(nowTime) || cronJob.nextTime.Equal(nowTime) {

					// 起一个 协程 执行调度任务
					go func(cronJobName string) {
						fmt.Printf("%v 被调度了, 时间：%v ", cronJobName, time.Now())
						fmt.Println("")
					}(cronJobName)

					cronJob.nextTime = cronJob.expr.Next(nowTime)
					// fmt.Println("next time: ", cronJob.nextTime)
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
