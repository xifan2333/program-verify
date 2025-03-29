FROM node AS builder-frontend

WORKDIR /frontend
# 复制 package.json 和 pnpm-lock.yaml
COPY ./frontend/package.json ./

# 安装依赖
RUN npm install

# 复制前端源代码
COPY ./frontend/ .

# 构建前端（使用环境变量）
ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=${VITE_API_BASE_URL:-/api/v1}
RUN npm run build

FROM golang AS builder-backend

WORKDIR /backend


ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux 
    
    # 复制 Go 依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 复制前端构建产物
COPY --from=builder-frontend /frontend/dist ./static

# 构建后端
RUN go build -ldflags="-s -w -extldflags '-static'" -o LicenseManager ./main.go

FROM alpine:latest

WORKDIR /app

# 复制后端二进制文件并设置权限
COPY --from=builder-backend /backend/LicenseManager /app/LicenseManager
RUN chmod +x /app/LicenseManager

# 复制前端文件
COPY --from=builder-frontend /frontend/dist /app/static

# 安装必要的系统包
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true

# 创建数据目录并设置权限
RUN mkdir -p /app/data && chmod 777 /app/data

EXPOSE 8080

# 使用JSON格式的CMD
CMD ["/app/LicenseManager"]


