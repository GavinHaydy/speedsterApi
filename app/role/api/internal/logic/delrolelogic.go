// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"
	"speedsterApi/app/role/api/internal/svc"
	"speedsterApi/app/role/api/internal/types"
	"speedsterApi/common/errno"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDelRoleLogic 删除角色
func NewDelRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRoleLogic {
	return &DelRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRoleLogic) DelRole(req *types.DelRoleReq) (resp *types.Response, err error) {
	count, err := l.svcCtx.SysRolePermission.FindByRoleId(l.ctx, req.Id)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("del role error, %v", err)
		return &types.Response{Code: errno.ErrSelectDbFailed}, err
	}
	if count > 0 {
		return &types.Response{Code: errno.ErrRoleNotDel}, errors.New("count > 0")
	}
	err = l.svcCtx.RoleModel.Delete(l.ctx, req.Id)
	if err != nil {
		return &types.Response{Code: errno.ErrDeleteFailed}, err
	}

	return &types.Response{Code: errno.Ok}, nil
}
