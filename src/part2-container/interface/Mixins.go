package main

import "fmt"

/*
- 接口灵活使用
- @author ch
- @date 2024/12/17 17:47
*/
type Selector struct {
	AWrite
}

type AWrite interface {
	write()
}

type DBWrite struct {
}

type CacheWrite struct {
}

func (db *DBWrite) write() {
	fmt.Println("db write")
}

func (db *CacheWrite) write() {
	fmt.Println("cache write")
}

func mainMutiInterFace() {
	var a AWrite = &Selector{
		&DBWrite{},
	}
	a.write()

	var b AWrite = &Selector{
		&CacheWrite{},
	}
	b.write()
}
