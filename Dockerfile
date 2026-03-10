FROM golang:1.20-alpine AS builder

# 使用国内镜像源加速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git build-base

# 设置Go模块代理
ENV GOPROXY=https://goproxy.cn,direct

# 复制go.mod文件
COPY go.mod .

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 生成go.sum文件
RUN go mod tidy

# 构建
RUN go build -o wechat-screenshot .

# 第二阶段，使用轻量级镜像
FROM alpine:latest

# 使用国内镜像源加速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装Chrome和必要的依赖
RUN apk add --no-cache chromium ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 安装网络工具
RUN apk add --no-cache curl wget

WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/wechat-screenshot /app/

# 设置入口点
ENTRYPOINT ["/app/wechat-screenshot"]
