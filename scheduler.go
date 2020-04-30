/******************************************************
# DESC       : 定时任务
# MAINTAINER : yamei
# EMAIL      : daixw@ecpark.cn
# DATE       : 2019/12/5
******************************************************/
package backend

import (
	"time"
)

// 定时任务
// do: 定时执行的任务
// interval:定时间隔
// after:关闭定时器后的清理工作
// return(chan):返回关闭定时器的开关通道
func Schedule(do func(), interval time.Duration, after ...func()) chan<- struct{} {
	ticker := time.NewTicker(interval)
	stopChan := make(chan struct{})
	SafetyRun(func() {
		defer ticker.Stop()
	LOOP:
		for {
			select {
			case <-ticker.C:
				do()
			case <-stopChan:
				if len(after) > 0 {
					for _, f := range after {
						f()
					}
				}
				close(stopChan)
				break LOOP
			}
		}
	})
	return stopChan
}

// 不强依赖系统的时钟
// 以传入时间为起始时间，每秒ticker一次
func Clock(startTime time.Time) (output <-chan time.Time, stop chan<- struct{}) {
	currentTimeChan := make(chan time.Time)
	return currentTimeChan, Schedule(func() {
		startTime = startTime.Add(1 * time.Second)
		currentTimeChan <- startTime
	}, 1*time.Second, func() {
		close(currentTimeChan)
	})
}
