package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁的作用就是将部分代码进行串行化，会导致性能下降，在读多写少的场景就应该使用读写锁

func mainReadWrite() {
	var num int
	var readWriteLock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		readWriteLock.Lock() //与其他写、读互斥
		num = 12
		fmt.Println("写")
		readWriteLock.Unlock()
	}()

	time.Sleep(time.Microsecond)

	go func() {
		defer wg.Done()
		readWriteLock.RLock() //读锁
		fmt.Println(num)
		fmt.Println("读")
		defer readWriteLock.RUnlock()
	}()

	wg.Wait()
}
