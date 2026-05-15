# speedsterApi

```shell
# api gen
# goctl api go -api user.api -dir .

#doc gen only once
# goctl env -w GOCTL_EXPERIMENTAL=on


goctl api swagger --api your.api --dir internal/handler/docs
#--api：指定 .api 文件路径
#--dir：输出目录
#--filename：指定生成的文件名（不含扩展名）
#--yaml：是否生成 YAML 格式

# dockerfile
goctl docker --go hello.go --port 8888

```

```
├── common/
│   ├── config/
│   ├── errors/
│   ├── utils/
│   └── response/
```