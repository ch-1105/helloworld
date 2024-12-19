package main

import (
	"fmt"
	"time"
)

// 线程间通信，不通过共享内存实现通信，而是通过通信来实现内存共享
/**
 * go中channel的应用场景:
	1.消息传递、消息过滤
	2.信号广播
	3.事件订阅和广播
	4.任务分发
	5.结果汇总
	6.并发控制
	7、同步和异步
*/
func mainCoroutine_communication() {
	// channel实现 channel分为有缓冲区和无缓冲区 也就是容量是否大于0
	// 有缓冲区channel 适用于生产者消费者场景
	// 无缓冲区channel 适用于第一时间通知任务完成
	var msg = make(chan string, 1) //如果定义不指定容量，会一直等待，会导致死锁，fatal error: all goroutines are asleep - deadlock!
	msg <- "hello"                 //将值放入channel
	msg1 := <-msg
	fmt.Println(msg1)

	// go的 happen before机制, 保证下列的chan赋值优先
	var happenBeforeMsg = make(chan string, 0)
	go func(happenBeforeMsg chan string) {
		data := <-happenBeforeMsg
		fmt.Println(data)
	}(happenBeforeMsg)

	happenBeforeMsg <- "hello2"
	time.Sleep(time.Microsecond * 1)

	// for range 取数据
	var forRangeMsg = make(chan string, 3)

	go func() {
		for data := range forRangeMsg {
			fmt.Println(data)
		}
		fmt.Println("msg done !")
	}()

	forRangeMsg <- "forRangeMsg hello3"
	forRangeMsg <- "forRangeMsg hello4"
	forRangeMsg <- "forRangeMsg hello5"

	close(forRangeMsg) //关闭后可以看到msg done !  已经关闭的channel 不能再放值 但可以取值

	//go func() {
	//	for data := range forRangeMsg {
	//		fmt.Println(data)
	//	}
	//	fmt.Println("msg done !")
	//}()

	//forRangeMsg <- "forRangeMsg channel close" //panic: send on closed channel

	time.Sleep(time.Microsecond * 1)

}
