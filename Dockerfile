# 后端构建阶段
FROM golang:1.23-alpine AS backend-builder

# 设置 Alpine 软件源为阿里云镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置 Go 代理为国内源
ENV GOPROXY=https://goproxy.cn,direct

# 设置 CGO 环境变量
ENV CGO_ENABLED=1

# 设置时区和语言
ENV TZ=Asia/Shanghai
ENV LANG=C.UTF-8

# 安装构建依赖
RUN apk add --no-cache \
    gcc \
    musl-dev \
    sqlite-dev \
    build-base

# 设置工作目录
WORKDIR /app

# 复制后端代码
COPY . .

# 下载依赖并构建
RUN go mod download
RUN go build -o main .
RUN chmod +x main

# 运行阶段
FROM alpine:latest

# 设置 Alpine 软件源为阿里云镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置时区和语言
ENV TZ=Asia/Shanghai
ENV LANG=C.UTF-8

# 安装运行时依赖
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    sqlite-dev \
    bash

# 设置工作目录
WORKDIR /app

# 从构建阶段复制必要文件
COPY --from=backend-builder /app/main .
COPY --from=backend-builder /app/static ./static
COPY --from=backend-builder /app/.env .

# 创建数据目录
RUN mkdir -p data && chmod 777 data

# 创建启动脚本
RUN echo '#!/bin/bash\n\
if [ "$1" = "shell" ]; then\n\
    exec /bin/bash\n\
else\n\
    exec ./main "$@"\n\
fi' > /app/entrypoint.sh && chmod +x /app/entrypoint.sh

# 暴露端口
EXPOSE 8080

# 设置入口点
ENTRYPOINT ["/app/entrypoint.sh"]

# 设置默认命令
CMD ["./main"]
