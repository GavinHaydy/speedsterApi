// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"speedsterApi/app/role/internal/config"
	"speedsterApi/app/role/model"
	"speedsterApi/common/middleware"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	RoleModel          model.RoleModel
	RedisJwtMiddleware rest.Middleware
	CasbinMiddleware   rest.Middleware
	DB                 sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DSN)
	rdb := redis.MustNewRedis(c.Redis)

	return &ServiceContext{
		Config:             c,
		RoleModel:          model.NewRoleModel(conn),
		RedisJwtMiddleware: middleware.NewRedisJwtMiddleware(rdb, c.Auth.AccessSecret).Handle,
		CasbinMiddleware:   middleware.NewCasbinMiddleware().Handle,
		DB:                 conn,
	}
}
