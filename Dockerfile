FROM golang:1.20

# 设置Go代理和时区
ENV GOPROXY=https://goproxy.io,direct
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 拷贝 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 拷贝源代码
COPY . .

# 编译应用
RUN go build -o main .

# 暴露 8080 端口
EXPOSE 8080

# 启动应用
CMD ["./main"]