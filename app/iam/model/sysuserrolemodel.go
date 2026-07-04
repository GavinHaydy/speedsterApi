package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserRoleModel = (*customSysUserRoleModel)(nil)

type (
	// SysUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserRoleModel.
	SysUserRoleModel interface {
		sysUserRoleModel
		withSession(session sqlx.Session) SysUserRoleModel
		FindRoleByUserId(ctx context.Context, userId string) (int64, error)
	}

	customSysUserRoleModel struct {
		*defaultSysUserRoleModel
	}
)

// NewSysUserRoleModel returns a model for the database table.
func NewSysUserRoleModel(conn sqlx.SqlConn) SysUserRoleModel {
	return &customSysUserRoleModel{
		defaultSysUserRoleModel: newSysUserRoleModel(conn),
	}
}

func (m *customSysUserRoleModel) withSession(session sqlx.Session) SysUserRoleModel {
	return NewSysUserRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSysUserRoleModel) FindRoleByUserId(ctx context.Context, userId string) (int64, error) {
	var resp SysUserRole
	query := fmt.Sprintf("select %s from %s where user_id = $1 ", sysUserRoleRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)

	switch {
	case err == nil:
		return resp.RoleId, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
