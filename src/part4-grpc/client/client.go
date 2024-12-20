package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"time"
)

type Res struct {
	Data int `json:"data"`
}

func add(a, b int) int {
	fmt.Println("开始运行")
	request := HttpRequest.NewRequest().SetTimeout(time.Millisecond * 100)
	str := "http://localhost:8080/add?a=" + fmt.Sprintf("%d", a) + "&b=" + fmt.Sprintf("%d", b)
	fmt.Println("str : ", str)
	response, err := request.Get(str)
	if err != nil {
		fmt.Println("err : ", err)
	}
	body, _ := response.Body()
	fmt.Println("body : ", string(body))
	res := Res{}
	_ = json.Unmarshal(body, &res)
	fmt.Println("结束运行")
	return res.Data
}

func main() {
	fmt.Println(add(192, 250))
}
