package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRolePermissionModel = (*customSysRolePermissionModel)(nil)

type (
	// SysRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRolePermissionModel.
	SysRolePermissionModel interface {
		sysRolePermissionModel
		withSession(session sqlx.Session) SysRolePermissionModel
		FindByRoleId(ctx context.Context, roleId int64) ([]int64, error)
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

func (m *customSysRolePermissionModel) FindByRoleId(ctx context.Context, roleId int64) (result []int64, err error) {
	var temp []*SysRolePermission
	query := fmt.Sprintf("select %s from %s where role_id = $1", sysRolePermissionRows, m.table)
	err = m.conn.QueryRowsCtx(ctx, &temp, query, roleId)
	switch {
	case err == nil:
		for _, v := range temp {
			result = append(result, v.PermissionId)
		}
		logx.Infof("=============%+v", result)
		return result, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}

}
