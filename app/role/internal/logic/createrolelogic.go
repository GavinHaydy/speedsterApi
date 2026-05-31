// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"speedsterApi/app/role/model"
	"speedsterApi/common/errno"

	"speedsterApi/app/role/internal/svc"
	"speedsterApi/app/role/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreateRoleLogic 新建角色
func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.NewRole) (resp *types.Response, err error) {
	insert, err := l.svcCtx.RoleModel.Insert(l.ctx, &model.Role{
		Name: req.Name,
		Code: req.Code,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
	})
	if err != nil {
		return &types.Response{Code: errno.ErrInsertFailed}, err
	}

	return &types.Response{Data: insert}, nil
}
