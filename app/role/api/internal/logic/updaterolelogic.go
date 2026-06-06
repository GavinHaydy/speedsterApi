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

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateRoleLogic 修改角色
func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRole) (resp *types.Response, err error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return &types.Response{Code: errno.ErrSelectDbFailed}, err
	}

	updateRole := &model.Role{
		Id:          role.Id,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		Status:      role.Status,
	}
	if req.Name != nil {
		updateRole.Name = *req.Name
	}
	if req.Code != nil {
		updateRole.Code = *req.Code
	}
	if req.Description != nil {
		updateRole.Description = sql.NullString{
			String: *req.Description,
			Valid:  true,
		}
	}
	if req.Status != nil {
		updateRole.Status = *req.Status
	}

	err = l.svcCtx.RoleModel.Update(l.ctx, updateRole)
	if err != nil {
		return &types.Response{Code: errno.ErrUpdateDataFailed}, err
	}

	return &types.Response{Code: errno.Ok}, nil
}
