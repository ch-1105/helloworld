package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Server is start")
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		fmt.Println("path : ", request.URL)
		a, _ := strconv.Atoi(request.Form["a"][0])
		b, _ := strconv.Atoi(request.Form["b"][0])
		writer.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(map[string]int{
			"result": a + b,
		})
		_, _ = writer.Write(data)
	})
	fmt.Println("Server is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
