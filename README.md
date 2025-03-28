# 软件授权管理系统

一个简单的软件授权管理系统，用于管理软件产品的许可证。

## 功能特性

- 管理员认证
- 产品管理
- 许可证生成和验证

## 技术栈

- Go
- Gin Web框架
- SQLite数据库
- JWT认证

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 初始化管理员账户

```bash
go run cmd/server/init_admin.go
```

默认管理员账户：

- 用户名：admin
- 密码：password

### 3. 启动服务器

```bash
go run main.go
```

服务器将在 <http://localhost:8080> 启动

## API接口

### 认证接口

```
POST /api/v1/auth/login
请求体：
{
    "username": "string",
    "password": "string"
}
```

### 产品接口

```
GET /api/v1/products
POST /api/v1/products
GET /api/v1/products/:id
```

### 许可证接口

```
POST /api/v1/licenses/generate
GET /api/v1/licenses
POST /api/v1/licenses/verify
```

## 配置

可以通过环境变量配置系统：

- JWT_SECRET: JWT密钥
- DB_PATH: 数据库文件路径
- PORT: 服务器端口

## 开发计划

- [ ] 批量生成许可证
- [ ] 许可证导出
- [ ] 使用统计
- [ ] 客户管理

## 配置说明

### 后端配置

1. 复制 `.env.example` 为 `.env`
2. 修改以下配置项：
   - `PORT`: 服务器端口号
   - `DB_PATH`: 数据库文件路径
   - `JWT_SECRET`: JWT密钥
   - `ADMIN_USERNAME`: 管理员用户名
   - `ADMIN_PASSWORD`: 管理员密码

### 前端配置

1. 进入前端目录：`cd frontend`
2. 复制 `.env.example` 为 `.env`
3. 修改以下配置项：
   - `VITE_API_BASE_URL`: API基础路径

## 开发环境设置

1. 安装依赖

   ```bash
   # 后端
   go mod download
   
   # 前端
   cd frontend
   pnpm install
   ```

2. 启动服务

   ```bash
   # 后端
   go run cmd/server/main.go
   
   # 前端
   cd frontend
   pnpm dev
   ```

## 生产环境部署

1. 构建前端

   ```bash
   cd frontend
   pnpm build
   ```

2. 构建后端

   ```bash
   go build -o server cmd/server/main.go
   ```

3. 运行服务

   ```bash
   ./server
   ```

## 注意事项

- 请确保在生产环境中修改所有默认密码和密钥
- 建议使用环境变量或配置文件管理敏感信息
- 定期备份数据库文件
