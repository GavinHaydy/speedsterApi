package svc

import (
	"speedsterApi/app/permission/model"
	"speedsterApi/app/permission/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config        config.Config
	SysPermission model.SysPermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DB.DSN)

	return &ServiceContext{
		Config:        c,
		SysPermission: model.NewSysPermissionModel(conn),
	}
}
