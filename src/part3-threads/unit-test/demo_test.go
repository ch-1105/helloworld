package main

import "testing"

func TestAdd(t *testing.T) {
	i := sum(1, 2)
	if i != 3 {
		t.Error("1+2=3, but", i)
	} else {
		t.Log("1+2=3")
	}
}

func TestAdd2(t *testing.T) {
	i := sum(1, 2)
	if testing.Short() {
		t.Skip("Skip") //根据情况，进行跳过执行
	}
	if i != 3 {
		t.Error("1+2=3, but", i)
	} else {
		t.Log("1+2=3 第二次")
	}
}
