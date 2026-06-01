# 开发环境搭建指南

本文档用于说明如何在本地搭建并运行 speedsterApi 开发环境。

---

## 开发方式

本项目支持两种开发方式：

### 方式一：本地安装依赖（推荐熟悉环境）

自行安装 PostgreSQL 和 Redis。

### 方式二：Docker 启动依赖（推荐快速启动）

使用 docker-compose 一键启动数据库依赖。

```bash id="runcompose"
docker compose -f docker-compose.dev.yml up -d
```

---

## 环境要求

开发前请确保已安装：

- Go
- PostgreSQL
- Redis

检查版本：

```bash id="checkver"
go version
psql --version
redis-server --version
```

---

## 数据库配置

项目默认使用 PostgreSQL。

### 创建数据库

```sql id="createdb"
CREATE DATABASE speedster;
```

### 导入初始化数据

执行初始化脚本：

```bash id="importsql"
psql -U <用户名> -d speedster -f common/init/init.sql
```

初始化脚本会创建以下表：

- sys_user
- role
- sys_user_role
- sys_permission
- sys_role_permission

并插入默认管理员账号及权限数据。

---

## Redis 配置

### 本地模式

```text id="redislocal"
Host: localhost:6379
Password: mypassword
```

### Docker 模式

```text id="redisdocker"
Host: localhost:16379
Password: mypassword
```

---

## PostgreSQL 配置

### 本地模式

```text id="pglocal"
Host: localhost:5432
Database: speedster
User: speedster
Password: speedster
```

### Docker 模式

```text id="pgdocker"
Host: localhost:15432
Database: speedster
User: speedster
Password: speedster
```

---

## 配置文件说明

### User 服务

```text id="userconf"
app/user/etc/user-api.yaml
```

默认端口：

```text id="userport"
8888
```

---

### Role 服务

```text id="roleconf"
app/role/etc/role-api.yaml
```

默认端口：

```text id="roleport"
8889
```

---

### Gateway 服务

```text id="gwconf"
app/gateway/etc/gateway-api.yaml
```

默认端口：

```text id="gwport"
9527
```

---

## 服务启动顺序

由于 User 和 Role 服务依赖数据库并加载权限策略，建议按以下顺序启动：

1. PostgreSQL
2. Redis
3. User 服务
4. Role 服务
5. Gateway 服务

---

## 启动服务

### 启动 User 服务

```bash id="runuser"
cd app/user
go run .
```

启动成功输出：

```text id="userok"
Starting server at 0.0.0.0:8888...
```

---

### 启动 Role 服务

```bash id="runrole"
cd app/role
go run .
```

启动成功输出：

```text id="roleok"
Starting server at 0.0.0.0:8889...
```

---

### 启动 Gateway 服务

```bash id="rungw"
cd app/gateway
go run .
```

启动成功输出：

```text id="gwok"
Starting server at 0.0.0.0:9527...
```

---

## 服务端口

| 服务       | 端口         |
| ---------- | ------------ |
| Gateway    | 9527         |
| User       | 8888         |
| Role       | 8889         |
| PostgreSQL | 5432 / 15432 |
| Redis      | 6379 / 16379 |

---

## API 文档生成

首次启用 Swagger：

```bash id="swaggerenv"
goctl env -w GOCTL_EXPERIMENTAL=on
```

生成 Swagger：

```bash id="swaggergen"
goctl api swagger \
  --api your.api \
  --dir internal/handler/docs
```

参数说明：

| 参数       | 说明         |
| ---------- | ------------ |
| --api      | API 文件路径 |
| --dir      | 输出目录     |
| --filename | 文件名       |
| --yaml     | 输出 YAML    |

---

## Go-Zero 代码生成

生成 API 代码：

```bash id="genapi"
goctl api go -api user.api -dir .
```

生成 Dockerfile：

```bash id="gendocker"
goctl docker --go hello.go --port 8888
```

---

## 常见问题

### PostgreSQL 连接失败

检查：

- PostgreSQL 是否启动
- 数据库是否存在
- DSN 是否正确（端口是否匹配 5432 / 15432）

---

### Redis 连接失败

检查：

- Redis 是否启动
- 密码是否匹配
- 端口是否正确（6379 / 16379）

---

### Casbin 初始化失败

检查：

- 是否执行 `common/init/init.sql`
- 权限表是否存在
- 数据库是否可连接
