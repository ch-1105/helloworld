package main

import (
	"fmt"
	"time"
)

/**
 * 交叉打印数字和字符
 */

var numSign, strSign = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		<-numSign
		i += 1
		fmt.Println(i)
		strSign <- true
		time.Sleep(time.Second * 1)
	}
}

func printStr() {
	for {
		<-strSign
		fmt.Println("a")
		numSign <- true
		time.Sleep(time.Second * 1)

	}
}

func Cross_printing() {
	go func() {
		printNum()
	}()
	go func() {
		printStr()
	}()
	numSign <- true
	time.Sleep(time.Second * 10)
}
