FROM node AS builder-frontend

WORKDIR /frontend
# 复制 package.json 和 pnpm-lock.yaml
COPY ./frontend/package.json ./

# 安装依赖
RUN npm install

# 复制前端源代码
COPY ./frontend/ .

# 构建前端（使用生产环境）
ENV NODE_ENV=production
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

# 复制源代码和前端构建产物
COPY . .
COPY --from=builder-frontend /frontend/dist ./frontend/dist

# 构建后端
RUN go build -ldflags="-s -w -extldflags '-static'" -o LicenseManager ./main.go

FROM alpine:latest

# 创建应用目录和数据目录
RUN mkdir -p /app/data \
    && chown -R nobody:nobody /app \
    && chmod -R 755 /app

WORKDIR /app

# 复制所有文件到应用目录
COPY --from=builder-backend /backend/LicenseManager /app/
COPY --from=builder-backend /backend/frontend/dist /app/frontend/dist

# 设置文件权限
RUN chown -R nobody:nobody /app \
    && chmod -R 755 /app \
    && chmod -R 777 /app/data

# 只声明数据目录为可挂载的卷
VOLUME ["/app/data"]

# 设置非root用户
USER nobody

EXPOSE 8080

# 使用JSON格式的CMD
CMD ["/app/LicenseManager"]


