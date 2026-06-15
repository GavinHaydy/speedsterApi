// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"speedsterApi/app/permission/api/internal/config"
	"speedsterApi/app/permission/api/internal/middleware"
	"speedsterApi/app/permission/rpc/pb"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	RedisJwtMiddleware rest.Middleware
	CasbinMiddleware   rest.Middleware
	PermissionRpc      pb.PermissionClient
	DB                 sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(c.PermissionRpc)
	conn := postgres.New(c.DB.DSN)

	return &ServiceContext{
		Config:             c,
		RedisJwtMiddleware: middleware.NewRedisJwtMiddleware().Handle,
		CasbinMiddleware:   middleware.NewCasbinMiddleware().Handle,
		PermissionRpc:      pb.NewPermissionClient(client.Conn()),
		DB:                 conn,
	}
}
