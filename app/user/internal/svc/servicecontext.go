// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"user/internal/config"
	"user/internal/middleware"
	"user/model"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	SysUserModel       model.SysUserModel
	Redis              redis.Redis
	RedisJwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DSN)
	//conn := sqlx.NewSqlConn(c.DB.Driver, c.DB.DSN)

	return &ServiceContext{
		Config:             c,
		SysUserModel:       model.NewSysUserModel(conn),
		Redis:              *redis.MustNewRedis(c.Redis),
		RedisJwtMiddleware: middleware.NewRedisJwtMiddleware(c, redis.MustNewRedis(c.Redis)).Handle,
	}
}
