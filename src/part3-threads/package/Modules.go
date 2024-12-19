package main

func main() {
	// 设置国内镜像代理
	_ = "https://goproxy.cn,direct" //当前版本默认，所以不用修改
	// 1.go env -w GO111MODULE=on
	// 2.go get ,go mod tidy ,go install
	// 3.go get -u 升级包 go get -u=patch 更新到修订版本
}
