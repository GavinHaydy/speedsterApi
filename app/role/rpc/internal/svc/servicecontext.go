package svc

import (
	"speedsterApi/app/role/model"
	"speedsterApi/app/role/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config           config.Config
	SysUserRoleModel model.SysUserRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DSN)
	return &ServiceContext{
		Config:           c,
		SysUserRoleModel: model.NewSysUserRoleModel(conn),
	}
}
