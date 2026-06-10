# Development Guide

## Environment Requirements

Before starting development, ensure the following software is installed:

| Software       | Recommended Version |
| -------------- | ------------------- |
| Go             | 1.26+               |
| Docker         | Latest              |
| Docker Compose | Latest              |
| Git            | Latest              |
| Make           | Latest              |
| tmux           | Latest (optional)   |
| Air            | Latest (optional)   |

Install Air:

```bash
go install github.com/air-verse/air@latest
```

---

## Clone Project

```bash
git clone <repository-url>
cd speedster
```

---

## Configuration

Create a local environment file:

```bash
cp .env.example .env
```

Example:

```env
POSTGRES_USER=speedster
POSTGRES_PASSWORD=password

POSTGRES_DB=speedster

REDIS_PASSWORD=password

JWT_SECRET=change_me
```

> Do not commit `.env` files to the repository.

---

## Project Structure

```text
speedster
├── app
│   ├── gateway
│   ├── user
│   │   ├── api
│   │   └── user
│   ├── role
│   │   ├── api
│   │   └── role
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

## Start Infrastructure

Start PostgreSQL and Redis:

```bash
make up
```

Check container status:

```bash
make ps
```

View logs:

```bash
make logs
```

Stop containers:

```bash
make down
```

---

## Local Development

Start all local services:

```bash
make start
```

Stop local services:

```bash
make stop
```

Restart local services:

```bash
make local-restart
```

The startup script uses tmux and Air for hot reload.

---

## Generate Swagger Documentation

Generate User API documentation:

```bash
goctl api plugin \
  -plugin goctl-swagger="swagger -filename user.json" \
  -api app/user/api/user.api \
  -dir app/user/api/docs
```

Generated file:

```text
app/user/api/docs/user.json
```

---

## Build Services

User API:

```bash
cd app/user/api

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -trimpath \
-ldflags="-s -w" \
-o user-api .
```

User RPC:

```bash
cd app/user/user

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -trimpath \
-ldflags="-s -w" \
-o user-rpc .
```

---

## Docker Deployment

Build and start containers:

```bash
make up
```

Rebuild:

```bash
make build
```

Restart:

```bash
make restart
```

Clean containers and volumes:

```bash
make clean
```

---

## Coding Conventions

### API

* API services provide HTTP interfaces.
* RPC services provide internal service communication.
* API services should not directly access databases.

### Configuration

* Do not commit passwords or secrets.
* Use environment variables for sensitive information.
* Commit only `.env.example`.

### Git

Feature branches:

```text
feature/xxx
```

Bugfix branches:

```text
fix/xxx
```

Commit examples:

```text
feat(user): add user list api
fix(role): resolve permission cache issue
refactor(common): optimize response structure
```

---

## Troubleshooting

### Docker Containers Not Starting

Check logs:

```bash
make logs
```

### PostgreSQL Connection Failed

Verify:

```bash
docker ps
```

and:

```bash
docker compose ps
```

### Redis Authentication Failed

Verify the `REDIS_PASSWORD` value in `.env`.

### Go Build Failed

Clean module cache:

```bash
go clean -modcache
go mod tidy
```

Then rebuild.

---

## Future Architecture

Planned microservice layout:

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

All inter-service communication should use go-zero RPC.
