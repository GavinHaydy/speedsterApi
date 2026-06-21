// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"speedsterApi/app/iam/api/internal/config"
	"speedsterApi/common/middleware"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	Redis              redis.Redis
	RedisJwtMiddleware rest.Middleware
	DB                 sqlx.SqlConn
	CasbinMiddleware   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	//client := zrpc.MustNewClient(c.UserRpc)
	conn := postgres.New(c.DB.DSN)
	//conn := sqlx.NewSqlConn(c.DB.Driver, c.DB.DSN)
	rdb := redis.MustNewRedis(c.Redis)

	return &ServiceContext{
		Config:             c,
		Redis:              *rdb,
		RedisJwtMiddleware: middleware.NewRedisJwtMiddleware(rdb, c.Auth.AccessSecret).Handle,
		DB:                 conn,
		CasbinMiddleware:   middleware.NewCasbinMiddleware().Handle,
	}
}
