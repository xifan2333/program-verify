# 前端构建阶段
FROM node:18-alpine AS frontend-builder
WORKDIR /app
COPY . .

WORKDIR /app/frontend
RUN npm install
# 创建前端环境变量文件
RUN echo "VITE_API_BASE_URL=${VITE_API_BASE_URL:-/api/v1}" > .env
RUN npm run build

# 后端构建阶段
FROM golang:1.23.4-alpine AS backend-builder
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 最终运行阶段
FROM alpine:latest
WORKDIR /app

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 创建后端环境变量文件
RUN echo "PORT=${PORT:-8080}\n\
DB_PATH=${DB_PATH:-data/program_verify.db}\n\
JWT_SECRET=${JWT_SECRET:-your-secret-key}\n\
ADMIN_USERNAME=${ADMIN_USERNAME:-admin}\n\
ADMIN_PASSWORD=${ADMIN_PASSWORD:-password}" > .env

RUN rm -rf /app/frontend

# 设置工作目录
WORKDIR /app/

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./main"]
