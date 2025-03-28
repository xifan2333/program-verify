FROM alpine:latest

# 设置 Alpine 软件源为阿里云镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装必要的构建工具和依赖
RUN apk add --no-cache \
    nodejs \
    npm \
    go \
    git \
    ca-certificates \
    tzdata \
    build-base \
    gcc \
    musl-dev \
    bash

# 设置 npm 镜像为淘宝镜像
RUN npm config set registry https://registry.npmmirror.com

# 设置 Go 代理为国内源
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制整个项目
COPY . .

# 构建前端
WORKDIR /app/frontend
RUN npm install
RUN echo "VITE_API_BASE_URL=${VITE_API_BASE_URL:-/api/v1}" > .env
RUN npm run build

# 构建后端
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 创建环境变量文件
RUN echo "PORT=${PORT:-8080}\n\
DB_PATH=${DB_PATH:-data/program_verify.db}\n\
JWT_SECRET=${JWT_SECRET:-your-secret-key}\n\
ADMIN_USERNAME=${ADMIN_USERNAME:-admin}\n\
ADMIN_PASSWORD=${ADMIN_PASSWORD:-password}" > .env

# 创建数据目录
RUN mkdir -p data

# 清理构建工具和依赖，但保留必要的 C 库和 bash
RUN apk del \
    nodejs \
    npm \
    go \
    git \
    build-base \
    gcc \
    musl-dev \
    && rm -rf /var/cache/apk/*

# 清理源码和构建文件，只保留必要文件
RUN rm -rf \
    /app/frontend/node_modules \
    /app/frontend/src \
    /app/frontend/public \
    /app/frontend/.vscode \
    /app/frontend/.env.example \
    /app/frontend/.env.production \
    /app/frontend/.npmrc \
    /app/frontend/tsconfig.* \
    /app/frontend/vite.config.ts \
    /app/frontend/uno.config.ts \
    /app/frontend/package.json \
    /app/frontend/pnpm-lock.yaml \
    /app/frontend/README.md \
    /app/internal \
    /app/go.mod \
    /app/go.sum \
    /app/.git \
    /app/.gitignore \
    /app/README.md \
    /app/api.http \
    /app/.env.example \
    /root/.npm \
    /root/.cache \
    /root/go \
    /tmp/*

# 设置工作目录
WORKDIR /app

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./main"]

# 设置入口为 bash
ENTRYPOINT ["/bin/bash"]
