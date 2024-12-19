package main

import (
	"fmt"
	"sync"
	"time"
)

func asyncPrint() {
	fmt.Println("打印中")
}
func coroutine() {
	// go是协程，go自身进行管理 不是操作系统管理的线程
	// 协程优势: 性能高，内存占用少，切换快，go只有协程
	for i := 0; i <= 10; i++ {
		go asyncPrint()
	}

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("匿名函数")
		}
	}()
	time.Sleep(time.Microsecond * 1)

	// 协程乱序的解决
	// 1. WaitGroup -> 用于等待协程执行
	var wg sync.WaitGroup
	wg.Add(50) //监控50个协程
	for i := 0; i < 50; i++ {
		//wg.Add(1)//或者在这里每次启动加一个
		go func(num int) {
			fmt.Println(num)
			defer wg.Done() // 减少一个协程,使用defer确保执行，放在函数内部第一位
		}(i)
	}
	// 等待
	wg.Wait()
}
