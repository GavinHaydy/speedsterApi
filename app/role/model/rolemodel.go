package model

import (
	"context"
	"fmt"
	"speedsterApi/app/role/internal/types"
	"speedsterApi/common/utils/timex"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		withSession(session sqlx.Session) RoleModel
		SelectRoleList(ctx context.Context, req *types.RoleListReq) (total int64, userList []*types.RoleListRsp, err error)
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn),
	}
}

func (m *customRoleModel) withSession(session sqlx.Session) RoleModel {
	return NewRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customRoleModel) SelectRoleList(ctx context.Context, req *types.RoleListReq) (total int64, userList []*types.RoleListRsp, err error) {
	var temp []*Role

	builder := squirrel.Select().From(m.table).PlaceholderFormat(squirrel.Dollar)
	if req.Rolename != "" {
		builder = builder.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", req.Rolename)})
	}
	if req.Status != 0 {
		builder = builder.Where(squirrel.Eq{"status": req.Status})
	}
	if req.Code != "" {
		builder = builder.Where(squirrel.Like{"code": fmt.Sprintf("%%%s%%", req.Code)})
	}

	countBuilder := builder.Columns("COUNT(*)").GroupBy("id")
	countQuery, countValues, err := countBuilder.ToSql()
	if err != nil {
		return 0, []*types.RoleListRsp{}, err
	}
	logx.Infof("========count sql:%s=========", countQuery)

	err = m.conn.QueryRowCtx(ctx, &total, countQuery, countValues...)
	if err != nil {
		return 0, []*types.RoleListRsp{}, err
	}

	if total == 0 {
		return 0, []*types.RoleListRsp{}, nil
	}

	// 添加分页和排序
	offset := (req.PageNo - 1) * req.PageSize
	query, values, err := builder.
		Columns(roleRows).
		Limit(uint64(req.PageSize)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return 0, []*types.RoleListRsp{}, err
	}

	err = m.conn.QueryRows(&temp, query, values...)
	if err != nil {
		logx.Infof("------------err:%v", err)
		return 0, []*types.RoleListRsp{}, err
	}

	rspList := make([]*types.RoleListRsp, 0, len(temp))
	for _, role := range temp {
		var tempDes string
		if role.Description.Valid {
			tempDes = role.Description.String
		}
		rspList = append(rspList, &types.RoleListRsp{
			Name:        role.Name,
			Code:        role.Code,
			Desctiption: tempDes,
			Status:      role.Status,
			CreateTime:  timex.Format(role.CreatedAt),
			UpdateTime:  timex.Format(role.UpdatedAt),
		})
	}
	return total, rspList, nil
}
