package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		withSession(session sqlx.Session) SysUserModel
		FindByAccountAndPW(ctx context.Context, account string, password string) (*SysUser, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *customSysUserModel) withSession(session sqlx.Session) SysUserModel {
	return NewSysUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSysUserModel) FindByAccountAndPW(ctx context.Context, username string, password string) (*SysUser, error) {
	var user SysUser
	// 1. 查询所有字段
	query := fmt.Sprintf("SELECT * FROM %s WHERE \"username\" = $1 AND \"password\" = $2 LIMIT 1", m.table)
	// 2. 将查询结果扫描到 user 结构体中
	// 注意：这里假设 QueryRowCtx 支持直接扫描到结构体，如果不支持，需要手动传入 &user.Id, &user.Account
	err := m.conn.QueryRowCtx(ctx, &user, query, username, password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
