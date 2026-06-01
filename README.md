# speedsterApi

基于 Go-Zero 的权限管理后端示例项目。

## Services

* Gateway (9527)
* User Service (8888)
* Role Service (8889)

## Development

开发环境搭建请参考：

[development.md](docs/development.md)

## Code Generation

### Generate API

```bash
goctl api go -api user.api -dir .
```

### Generate Swagger

```bash
goctl env -w GOCTL_EXPERIMENTAL=on

goctl api swagger \
  --api your.api \
  --dir internal/handler/docs
```

### Generate Dockerfile

```bash
goctl docker --go hello.go --port 8888
```
