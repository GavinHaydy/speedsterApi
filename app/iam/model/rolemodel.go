package model

import (
	"context"
	"fmt"
	"speedsterApi/app/iam/rpc/pb"
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
		SelectRoleList(ctx context.Context, req *pb.RoleListReq) (roleList *pb.RoleListResp, err error)
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

func (m *customRoleModel) SelectRoleList(ctx context.Context, req *pb.RoleListReq) (roleList *pb.RoleListResp, err error) {
	var temp []*Role
	var total int64

	builder := squirrel.Select().From(m.table).PlaceholderFormat(squirrel.Dollar)
	if req.RoleName != "" {
		builder = builder.Where(squirrel.Like{"name": fmt.Sprintf("%%%s%%", req.RoleName)})
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
		return nil, err
	}
	logx.Infof("========count sql:%s=========", countQuery)

	err = m.conn.QueryRowCtx(ctx, &total, countQuery, countValues...)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, nil
	}

	// 添加分页和排序
	offset := (req.PageNo - 1) * req.PageSize
	query, values, err := builder.
		Columns(roleRows).
		Limit(uint64(req.PageSize)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = m.conn.QueryRows(&temp, query, values...)
	if err != nil {
		logx.Infof("------------err:%v", err)
		return nil, err
	}

	rspList := make([]*pb.RoleListItem, 0, len(temp))
	for _, role := range temp {
		var tempDes string
		if role.Description.Valid {
			tempDes = role.Description.String
		}
		rspList = append(rspList, &pb.RoleListItem{
			Name:        role.Name,
			Code:        role.Code,
			Description: &tempDes,
			Status:      role.Status,
			CreateTime:  timex.Format(role.CreatedAt),
			UpdateTime:  timex.Format(role.UpdatedAt),
		})
	}
	return &pb.RoleListResp{List: rspList, Total: total}, nil
}
