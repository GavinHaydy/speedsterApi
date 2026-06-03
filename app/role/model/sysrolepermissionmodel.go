package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRolePermissionModel = (*customSysRolePermissionModel)(nil)

type (
	// SysRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRolePermissionModel.
	SysRolePermissionModel interface {
		sysRolePermissionModel
		withSession(session sqlx.Session) SysRolePermissionModel
		FindByRoleId(ctx context.Context, roleId int64) (int, error)
	}

	customSysRolePermissionModel struct {
		*defaultSysRolePermissionModel
	}
)

// NewSysRolePermissionModel returns a model for the database table.
func NewSysRolePermissionModel(conn sqlx.SqlConn) SysRolePermissionModel {
	return &customSysRolePermissionModel{
		defaultSysRolePermissionModel: newSysRolePermissionModel(conn),
	}
}

func (m *customSysRolePermissionModel) withSession(session sqlx.Session) SysRolePermissionModel {
	return NewSysRolePermissionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSysRolePermissionModel) FindByRoleId(ctx context.Context, roleId int64) (int, error) {
	var count int

	query := `
		SELECT COUNT(*)
		FROM sys_role_permission
		WHERE role_id = $1	
	`

	err := m.conn.QueryRowCtx(ctx, &count, query, roleId)
	if err != nil {
		return 0, err
	}

	return count, nil
}
