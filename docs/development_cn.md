# 开发指南

## 环境要求

开始开发前，请确保已安装以下软件：

| 软件             | 推荐版本    |
| -------------- | ------- |
| Go             | 1.26+   |
| Docker         | 最新版     |
| Docker Compose | 最新版     |
| Git            | 最新版     |
| Make           | 最新版     |
| tmux           | 最新版（可选） |
| Air            | 最新版（可选） |

安装 Air：

```bash
go install github.com/air-verse/air@latest
```

---

## 获取项目

```bash
git clone <repository-url>
cd speedster
```

---

## 配置环境变量

复制示例配置文件：

```bash
cp .env.example .env
```

示例：

```env
POSTGRES_USER=speedster
POSTGRES_PASSWORD=password

POSTGRES_DB=speedster

REDIS_PASSWORD=password

JWT_SECRET=change_me
```

> ⚠️ 请勿将 `.env` 提交到仓库。

---

## 项目结构

```text
speedster
├── app
│   ├── gateway
│   │
│   ├── user
│   │   ├── api
│   │   └── user
│   │
│   ├── role
│   │   ├── api
│   │   └── role
│   │
│   └── permission
│       ├── api
│       └── permission
│
├── common
├── deploy
│   ├── docker-compose.yml
│   └── *.Dockerfile
│
├── scripts
│   └── start.sh
│
├── docs
└── Makefile
```

---

## 启动基础设施

启动 PostgreSQL 和 Redis：

```bash
make up
```

查看容器状态：

```bash
make ps
```

查看日志：

```bash
make logs
```

停止容器：

```bash
make down
```

---

## 本地开发

启动本地所有服务：

```bash
make start
```

停止本地服务：

```bash
make stop
```

重启本地服务：

```bash
make local-restart
```

项目使用：

* tmux 管理多个服务窗口
* Air 实现热重载

---

## 生成 Swagger 文档

以 User 服务为例：

```bash
goctl api plugin \
  -plugin goctl-swagger="swagger -filename user.json" \
  -api app/user/api/user.api \
  -dir app/user/api/docs
```

生成文件：

```text
app/user/api/docs/user.json
```

---

## 编译服务

### User API

```bash
cd app/user/api

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -trimpath \
-ldflags="-s -w" \
-o user-api .
```

### User RPC

```bash
cd app/user/user

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -trimpath \
-ldflags="-s -w" \
-o user-rpc .
```

---

## Docker 部署

构建并启动：

```bash
make up
```

重新构建镜像：

```bash
make build
```

重启容器：

```bash
make restart
```

清理容器及数据卷：

```bash
make clean
```

---

## 开发规范

### 服务职责

#### API 服务

负责：

* HTTP 接口
* 参数校验
* 用户认证
* 权限校验

不负责：

* 数据库直接访问
* 业务逻辑实现

API 应通过 RPC 调用对应服务。

---

#### RPC 服务

负责：

* 业务逻辑处理
* 数据库访问
* 缓存操作
* 服务间调用

---

### 配置规范

敏感信息禁止提交到仓库：

* 数据库密码
* Redis 密码
* JWT Secret
* 第三方平台密钥

统一使用环境变量配置。

仓库中仅保留：

```text
.env.example
```

---

### Git 分支规范

功能开发：

```text
feature/xxx
```

Bug 修复：

```text
fix/xxx
```

重构：

```text
refactor/xxx
```

---

### Commit 规范

示例：

```text
feat(user): 新增用户列表接口
fix(role): 修复角色权限缓存问题
refactor(common): 重构统一响应结构
docs: 更新开发文档
```

---

## 常见问题

### Docker 无法启动

查看日志：

```bash
make logs
```

查看容器状态：

```bash
make ps
```

---

### PostgreSQL 连接失败

检查数据库是否启动：

```bash
docker ps
```

检查配置：

```bash
cat .env
```

---

### Redis 密码错误

检查：

```env
REDIS_PASSWORD=xxxxxx
```

是否与 Docker Compose 配置一致。

---

### Go 编译失败

尝试清理依赖缓存：

```bash
go clean -modcache
go mod tidy
```

然后重新编译。

---

## 未来规划

项目采用 go-zero 微服务架构。

规划中的服务结构：

```text
gateway

user-api
user-rpc

role-api
role-rpc

permission-api
permission-rpc

order-api
order-rpc

etcd

postgres
redis
```

服务之间统一通过 RPC 通信。

Gateway 负责：

* API 聚合
* 身份认证
* 权限控制
* 路由转发

后续将逐步完善服务治理、链路追踪、监控告警等能力。
