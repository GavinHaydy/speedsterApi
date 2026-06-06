// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"speedsterApi/app/user/api/internal/config"
	"speedsterApi/app/user/model"
	"speedsterApi/app/user/user/pb/pb"
	"speedsterApi/common/middleware"

	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	SysUserModel       model.SysUserModel
	SysRoleModel       model.RoleModel
	SysUserRoleModel   model.SysUserRoleModel
	Redis              redis.Redis
	RedisJwtMiddleware rest.Middleware
	DB                 sqlx.SqlConn
	CasbinMiddleware   rest.Middleware
	UserRpc            pb.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(c.UserRpc)
	conn := postgres.New(c.DB.DSN)
	//conn := sqlx.NewSqlConn(c.DB.Driver, c.DB.DSN)
	rdb := redis.MustNewRedis(c.Redis)

	return &ServiceContext{
		Config:             c,
		SysUserModel:       model.NewSysUserModel(conn),
		SysRoleModel:       model.NewRoleModel(conn),
		SysUserRoleModel:   model.NewSysUserRoleModel(conn),
		Redis:              *rdb,
		RedisJwtMiddleware: middleware.NewRedisJwtMiddleware(rdb, c.Auth.AccessSecret).Handle,
		DB:                 conn,
		CasbinMiddleware:   middleware.NewCasbinMiddleware().Handle,
		UserRpc:            pb.NewUserClient(client.Conn()),
	}
}
