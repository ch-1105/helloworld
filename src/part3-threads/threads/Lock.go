package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var total int
var lock sync.Mutex  // 使用锁不能进行复制，复制后失去效果
var num atomic.Int32 // 原子化

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//total += 1
		num.Add(1)
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//total -= 1
		num.Add(-1)
		//lock.Unlock()
	}
}
func Lock() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total) //不加锁的情况下每次结果不同
}
