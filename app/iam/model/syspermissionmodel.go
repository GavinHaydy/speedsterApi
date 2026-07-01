package model

import (
	"context"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPermissionModel = (*customSysPermissionModel)(nil)

type (
	// SysPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPermissionModel.
	SysPermissionModel interface {
		sysPermissionModel
		withSession(session sqlx.Session) SysPermissionModel
		SelectTree(ctx context.Context, req *pb.PermissionTreeReq) ([]*SysPermission, error)
	}

	customSysPermissionModel struct {
		*defaultSysPermissionModel
	}
)

// NewSysPermissionModel returns a model for the database table.
func NewSysPermissionModel(conn sqlx.SqlConn) SysPermissionModel {
	return &customSysPermissionModel{
		defaultSysPermissionModel: newSysPermissionModel(conn),
	}
}

func (m *customSysPermissionModel) withSession(session sqlx.Session) SysPermissionModel {
	return NewSysPermissionModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSysPermissionModel) SelectTree(ctx context.Context, req *pb.PermissionTreeReq) ([]*SysPermission, error) {

	builder := squirrel.
		Select("*").
		From(m.table).
		OrderBy("sort asc", "id asc").
		PlaceholderFormat(squirrel.Dollar)

	if req.Name != "" {
		builder = builder.Where(
			squirrel.Like{
				"name": "%" + req.Name + "%",
			},
		)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var list []*SysPermission

	err = m.conn.QueryRowsCtx(
		ctx,
		&list,
		query,
		args...,
	)

	return list, err
}
