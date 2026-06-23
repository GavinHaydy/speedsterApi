package svc

import (
	"speedsterApi/app/iam/model"
	"speedsterApi/app/iam/rpc/internal/config"
	"speedsterApi/common/middleware"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	SysUserModel       model.SysUserModel
	SysRoleModel       model.RoleModel
	SysUserRoleModel   model.SysUserRoleModel
	Redis              redis.Redis
	RedisJwtMiddleware rest.Middleware
	DB                 sqlx.SqlConn
	CasBinMiddleware   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DSN)
	rdb := redis.MustNewRedis(c.CacheRedis)

	return &ServiceContext{
		Config:             c,
		SysUserModel:       model.NewSysUserModel(conn),
		SysRoleModel:       model.NewRoleModel(conn),
		SysUserRoleModel:   model.NewSysUserRoleModel(conn),
		Redis:              *rdb,
		RedisJwtMiddleware: middleware.NewRedisJwtMiddleware(rdb, c.CacheAuth.AccessSecret).Handle,
		DB:                 conn,
		CasBinMiddleware:   middleware.NewCasbinMiddleware().Handle,
	}
}
