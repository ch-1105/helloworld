package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

// 使用chan传值select实现监控
func cpuInfo(sign chan struct{}) {
	defer waitGroup.Done()
	for {
		select {
		case <-sign:
			fmt.Println("cpu info monitor over")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("cpu info")
		}
	}
}

// 使用context 实现监控
func cpuInfoByContext(sign context.Context) {
	defer waitGroup.Done()
	for {
		select {
		case <-sign.Done():
			fmt.Println("cpu Info By Context monitor over")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("cpu Info By Context")
		}
	}
}
func main() {
	// chan传值实现监控
	var sign = make(chan struct{})
	waitGroup.Add(1)
	go cpuInfo(sign)
	time.Sleep(time.Second * 5)
	sign <- struct{}{}
	waitGroup.Wait()
	fmt.Println("监控结束")

	// context实现监控
	waitGroup.Add(1)
	// 返回三种不同类型的结构体 withCancel withTimeout 按时间自动取消 withValue 传递全局链路信息
	// 这里的cancel具有传递性
	var ctx, cancel = context.WithCancel(context.Background())

	go cpuInfoByContext(ctx)
	time.Sleep(time.Second * 5)

	cancel()
	waitGroup.Wait()

}
