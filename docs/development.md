# 开发环境搭建指南

本文档用于说明如何在本地搭建并运行 speedsterApi 开发环境。

## 环境要求

开发前请确保已安装以下软件：

- Go
- PostgreSQL
- Redis

检查版本：

```bash
go version
psql --version
redis-server --version
```

---

## 数据库配置

项目默认使用 PostgreSQL。

### 创建数据库

```sql
CREATE DATABASE speedster;
```

### 导入初始化数据

执行项目中的初始化脚本：

```bash
psql -U <用户名> -d speedster -f common/init/init.sql
```

初始化脚本会自动创建：

- sys_user
- role
- sys_user_role
- sys_permission
- sys_role_permission

并插入默认管理员账号及权限数据。

---

## Redis 配置

项目默认配置：

```text
Host: redis-db:6379
Password: mypassword
```

可使用 Docker 快速启动：

```bash
docker run -d \
  --name redis-db \
  -p 6379:6379 \
  redis:7 \
  redis-server --requirepass mypassword
```

如使用本地 Redis，请根据实际情况修改配置文件。

---

## 配置文件说明

### User 服务

配置文件：

```text
app/user/etc/user-api.yaml
```

默认端口：

```text
8888
```

---

### Role 服务

配置文件：

```text
app/role/etc/role-api.yaml
```

默认端口：

```text
8889
```

---

### Gateway 服务

配置文件：

```text
app/gateway/etc/gateway-api.yaml
```

默认端口：

```text
9527
```

---

## 服务启动顺序

由于 User 和 Role 服务启动时需要连接 PostgreSQL 并加载权限策略，因此建议按以下顺序启动：

1. PostgreSQL
2. Redis
3. User 服务
4. Role 服务
5. Gateway 服务

---

## 启动服务

### 启动 User 服务

```bash
cd app/user
go run .
```

正常启动后会看到：

```text
Starting server at 0.0.0.0:8888...
```

---

### 启动 Role 服务

```bash
cd app/role
go run .
```

正常启动后会看到：

```text
Starting server at 0.0.0.0:8889...
```

---

### 启动 Gateway 服务

```bash
cd app/gateway
go run .
```

正常启动后会看到：

```text
Starting server at 0.0.0.0:9527...
```

---

## 服务端口

| 服务       | 端口 |
| ---------- | ---- |
| Gateway    | 9527 |
| User       | 8888 |
| Role       | 8889 |
| PostgreSQL | 5432 |
| Redis      | 6379 |

---

## API 文档生成

首次使用 Swagger 功能时：

```bash
goctl env -w GOCTL_EXPERIMENTAL=on
```

生成 Swagger：

```bash
goctl api swagger \
  --api your.api \
  --dir internal/handler/docs
```

参数说明：

| 参数       | 说明           |
| ---------- | -------------- |
| --api      | API 文件路径   |
| --dir      | 输出目录       |
| --filename | 输出文件名     |
| --yaml     | 生成 YAML 格式 |

---

## Go-Zero 代码生成

生成 API 代码：

```bash
goctl api go -api user.api -dir .
```

生成 Dockerfile：

```bash
goctl docker --go hello.go --port 8888
```

---

## 常见问题

### PostgreSQL 连接失败

请检查：

- PostgreSQL 是否启动
- 数据库是否已创建
- 配置文件中的 DSN 是否正确

---

### Redis 连接失败

请检查：

- Redis 是否启动
- Redis 密码是否与配置文件一致
- Redis 地址是否可访问

---

### Casbin 初始化失败

请检查：

- 是否执行了 `common/init/init.sql`
- 权限相关数据表是否存在
- PostgreSQL 是否正常连接
