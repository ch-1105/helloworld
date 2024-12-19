package StructDemo

import (
	"errors"
	"fmt"
)

func testErr() (a int, err error) {
	return a, errors.New("错误测试")
}
func main3() {
	// 函数错误处理
	// 错误处理理念: try catch => 返回err变量
	// go中必须处理错误 所以会出现大量 if err != nil
	_, err := testErr()
	if err != nil {
		fmt.Println(err)
	}

	// panic 类似java的Err 直接导致程序退出
	//panic("panic")
	// 执行不通
	//fmt.Println("main")

	// 进行捕获panic 此时panic 不会导致程序退出 其实就是个try catch, go这种用法用得很少
	defer func() { // 放在panic前
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic")

}
