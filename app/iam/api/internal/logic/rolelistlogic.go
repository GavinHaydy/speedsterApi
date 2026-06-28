// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/iam/rpc/iam"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

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
	list, err := l.svcCtx.IamRpc.RoleList(l.ctx, &iam.RoleListReq{
		RoleName: req.Rolename,
		Code:     req.Code,
		Status:   req.Status,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{Code: code, Msg: msg}, err
	}

	return &types.Response{Data: list}, nil
}
