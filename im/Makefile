.PHONY: build run clean

# 构建
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o websocket main.go

# 运行
run:
	@go run ./

# 代码格式化
format:
	go fmt ./

# 代码检测
check:
	go vet ./

# 清除二进制
clean:
	@if [ -f ${BIN_NAME} ] ; then rm ${BIN_NAME} ; fi

# 构建Docker镜像
image:
	docker build -t lab/${BIN_NAME}:latest .