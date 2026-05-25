// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/common/errno"

	"speedsterApi/app/role/internal/svc"
	"speedsterApi/app/role/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRoleListLogic 角色列表
func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	total, list, err := l.svcCtx.RoleModel.SelectRoleList(l.ctx, req)
	if err != nil {
		return &types.Response{Code: errno.ErrSelectDbFailed}, err
	}
	result := map[string]interface{}{
		"list":  list,
		"total": total,
	}
	return &types.Response{Data: result}, nil
}
