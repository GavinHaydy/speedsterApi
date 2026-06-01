# 开发环境搭建指南

本文档用于说明如何在本地搭建并运行 speedsterApi 开发环境。

---

## 开发方式

本项目支持两种开发方式：

### 方式一：本地运行应用服务

本地运行 User、Role 和 Gateway，PostgreSQL / Redis 需要自行安装并启动。

```bash id="runlocal"
make local-dev
```

停止本地服务：

```bash id="stoplocal"
make local-stop
```

查看本地日志：

```bash id="logslocal"
make local-logs
```

### 方式二：容器启动完整开发环境（推荐快速启动）

使用 `make dev` 子命令启动 PostgreSQL、Redis、User、Role 和 Gateway。Makefile 会优先使用 Podman，未安装 Podman 时自动使用 Docker；两者都不存在时退出。

```bash id="runcompose"
make dev up
```

常用子命令：

```bash id="composecommands"
make dev build   # 构建镜像
make dev up      # 构建并启动
make dev ps      # 查看容器状态
make dev logs    # 查看容器日志
make dev stop    # 停止容器
make dev clean   # 停止并删除数据卷
```

等价底层命令：

```bash id="runcomposefull"
podman compose -f docker-compose.dev.yml up -d --build
# 或
docker compose -f docker-compose.dev.yml up -d --build
```

---

## 环境要求

本地运行应用服务需要：

- Go
- PostgreSQL
- Redis

容器运行完整开发环境需要：

- Podman / Podman Compose 或 Docker / Docker Compose

检查版本：

```bash id="checkver"
go version
psql --version
redis-server --version
podman --version # 或 docker --version
```

---

## 数据库配置

项目默认使用 PostgreSQL。

### 本地模式

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

### 容器模式

容器模式会在 PostgreSQL 数据卷首次创建时自动执行：

```text id="containerinit"
common/init/init.sql
```

如果已经创建过数据卷，初始化脚本不会重复执行。需要重建数据卷时使用：

```bash id="composeclean"
make dev clean
make dev up
```

---

## Redis 配置

### 本地模式

```text id="redislocal"
Host: localhost:6379
Password: 空
```

### 容器模式

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

### 容器模式

```text id="pgdocker"
Host: localhost:15432
Database: speedster
User: speedster
Password: speedster
```

---

## 配置文件说明

### User 服务

本地运行使用：

```text id="userconf"
app/user/etc/user-api.local.yaml
```

容器启动使用：

```text id="userconfdocker"
app/user/etc/user-api.docker.yaml
```

默认端口：

```text id="userport"
8888
```

---

### Role 服务

本地运行使用：

```text id="roleconf"
app/role/etc/role-api.local.yaml
```

容器启动使用：

```text id="roleconfdocker"
app/role/etc/role-api.docker.yaml
```

默认端口：

```text id="roleport"
8889
```

---

### Gateway 服务

本地运行使用：

```text id="gwconf"
app/gateway/etc/gateway-api.local.yaml
```

容器启动使用：

```text id="gwconfdocker"
app/gateway/etc/gateway-api.docker.yaml
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

### 使用 Makefile 启动本地服务

```bash id="makelocal"
make local-dev
```

`make local-dev` 会分别启动：

- User：`app/user/etc/user-api.local.yaml`
- Role：`app/role/etc/role-api.local.yaml`
- Gateway：`app/gateway/etc/gateway-api.local.yaml`

### 单独启动 User 服务

```bash id="runuser"
cd app/user
go run . -f etc/user-api.local.yaml
```

启动成功输出：

```text id="userok"
Starting server at 0.0.0.0:8888...
```

---

### 单独启动 Role 服务

```bash id="runrole"
cd app/role
go run . -f etc/role-api.local.yaml
```

启动成功输出：

```text id="roleok"
Starting server at 0.0.0.0:8889...
```

---

### 单独启动 Gateway 服务

```bash id="rungw"
cd app/gateway
go run . -f etc/gateway-api.local.yaml
```

启动成功输出：

```text id="gwok"
Starting server at 0.0.0.0:9527...
```

---

## 服务端口

本地运行默认端口：

| 服务       | 地址           |
| ---------- | -------------- |
| Gateway    | localhost:9527 |
| User       | localhost:8888 |
| Role       | localhost:8889 |
| PostgreSQL | localhost:5432 |
| Redis      | localhost:6379 |

容器开发环境默认暴露到宿主机：

| 服务       | 宿主机地址      | 容器内地址         |
| ---------- | --------------- | ------------------ |
| Gateway    | localhost:9527  | gateway:9527       |
| PostgreSQL | localhost:15432 | postgres:5432      |
| Redis      | localhost:16379 | redis:6379         |
| User       | 不暴露          | user-service:8888  |
| Role       | 不暴露          | role-service:8889  |

可通过环境变量覆盖端口：

```bash id="composeports"
GATEWAY_PORT=19527 POSTGRES_PORT=25432 REDIS_PORT=26379 make dev up
```

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
- 本地模式 DSN 是否使用 `localhost:5432`
- 容器模式 DSN 是否使用 `postgres:5432`
- 如果使用 PostgreSQL 18 镜像并反复退出，确认 compose 挂载路径是 `/var/lib/postgresql`

---

### Redis 连接失败

检查：

- Redis 是否启动
- 密码是否匹配
- 本地模式 Redis 使用 `localhost:6379` 且密码为空
- 容器模式 Redis 使用 `redis:6379` 且密码为 `mypassword`

---

### Casbin 初始化失败

检查：

- 是否执行 `common/init/init.sql`
- 权限表是否存在
- 数据库是否可连接
- 容器镜像内是否包含 `/common/casbin/model.conf`

### Gateway 代理到 127.0.0.1:\<port\>

如果日志出现 `proxyconnect tcp: dial tcp 127.0.0.1:<port>: connect: connection refused`，说明容器可能继承了宿主机代理环境变量。开发 compose 已显式清空 app 容器中的 `HTTP_PROXY` / `HTTPS_PROXY`，并设置 `NO_PROXY=*`，修改后需要重新创建容器：

```bash id="composeproxyfix"
make dev stop
make dev up
```
