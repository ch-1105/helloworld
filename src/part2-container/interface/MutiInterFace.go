package main

import "fmt"

type A interface {
	haha()
}
type B interface {
	xixi()
}
type C struct {
	A
}
type D struct {
	A
}

func (c *C) haha() {
	fmt.Println("c -> haha")
}

func (c *D) haha() {
	fmt.Println("d -> haha")
}

func (c C) xixi() {
	fmt.Println("xixi")
}

func main() {
	// 这里接口接收的是C类型 而不是C指针，所以使用方式与InterFace.go文件不同
	var a A = &C{}
	a.haha()
	var b B = C{}
	b.xixi()
	var aa A = &D{}
	aa.haha()
}
