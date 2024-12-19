package container

import (
	"container/list"
	"fmt"
)

/*
 * 链表
 * @author ch
 * @date 2024/12/16 15:36
 * @param null
 */
func main5() {
	var list2 = list.New()
	list2.PushFront("avc") //头部
	list2.PushBack("bvc")  //尾部放
	list2.PushBack("cvc")
	list2.InsertBefore("121", list2.Front().Next()) //在avc后 bvc前插入
	for e := list2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	var myList list.List

	myList.PushBack("1")
	myList.PushBack("2")
	myList.PushBack("3")

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := myList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
