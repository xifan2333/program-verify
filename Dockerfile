FROM alpine:latest

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
    musl-dev

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

# 清理构建工具和依赖
RUN apk del \
    nodejs \
    npm \
    go \
    git \
    build-base \
    gcc \
    musl-dev

# 设置工作目录
WORKDIR /app

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./main"]
