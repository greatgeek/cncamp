# 使用官方的 Golang 镜像作为构建基础
FROM golang:1.17-alpine AS builder

# 在容器中设置工作目录
WORKDIR /app

# 把模块复制到容器中
COPY src/httpServer/go.mod ./

# 下载所有依赖
RUN go mod download

# 复制其它文件
COPY src/httpServer/. .

# 为 Go 应用程序构建一个可执行文件
RUN go build -o main .

# 使用 alpine 作为运行时镜像
FROM alpine:latest

# 从 builder 镜像中把 /app/main/ 复制到当前目录的 /main
COPY --from=builder /app/main /main

# 启动应用程序
ENTRYPOINT ["/main"]
