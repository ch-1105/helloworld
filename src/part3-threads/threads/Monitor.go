package main

import (
	"fmt"
	"time"
)

/**
 * @author ch
 * @date 2024/12/19 16:20
 * 监控两个协程执行
 */

var done = make(chan struct{})

func t1(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
}

func t2(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
}

func Monitor() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
		t1(ch1)
	}()

	go func() {
		t2(ch2)
	}()
	timer := time.NewTimer(time.Second * 2)
	for {
		select { // 某分支就绪了就执行哪个，当两个同时执行 随机进行选择
		case <-ch1:
			fmt.Println("ch1 done")
		case <-ch2:
			fmt.Println("ch2 done")
		case <-timer.C: // 超时后自动结束
			fmt.Println("timeout")
			return
		}
	}

}
