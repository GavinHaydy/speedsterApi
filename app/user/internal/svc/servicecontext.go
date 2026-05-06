// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"user/internal/config"
	"user/model"

	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config       config.Config
	SysUserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DSN)
	//conn := sqlx.NewSqlConn(c.DB.Driver, c.DB.DSN)

	return &ServiceContext{
		Config:       c,
		SysUserModel: model.NewSysUserModel(conn),
	}
}
