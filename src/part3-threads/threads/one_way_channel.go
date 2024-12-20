package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
* 单项channel 记录到3-10
* @author ch
* @date 2024/12/18 16:30
 */

func producer(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- strconv.Itoa(i * i)
	}
}

func consumer(c <-chan string) {
	for data := range c {
		fmt.Println(data)
	}
}
func one_way_channel() {
	c := make(chan string, 3)
	for i := 0; i < 3; i++ {
		c <- strconv.Itoa(i)
		fmt.Println(<-c)
	}
	time.Sleep(time.Microsecond * 1)
	//只能写入的channel
	//var oneWayWriteChannel chan<- string = c
	//只能读的channel
	//var oneWayReadChannel <-chan string = c
	//oneWayWriteChannel <- "hello"
	//<- oneWayWriteChannel | Invalid operation: <- oneWayWriteChannel (receive from the send-only type chan<- string)
	//oneWayReadChannel <-  | Invalid operation: oneWayReadChannel <- (send to the receive-only type <-chan string)
	time.Sleep(time.Microsecond * 1)

	go func() {
		producer(c)
	}()

	go func() {
		consumer(c)
	}()

	time.Sleep(time.Microsecond * 1)
}
