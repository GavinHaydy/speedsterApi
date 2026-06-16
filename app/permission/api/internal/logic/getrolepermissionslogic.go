// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/permission/rpc/pb"
	"speedsterApi/common/errno"

	"speedsterApi/app/permission/api/internal/svc"
	"speedsterApi/app/permission/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetRolePermissionsLogic 角色权限
func NewGetRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionsLogic {
	return &GetRolePermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermissionsLogic) GetRolePermissions(req *types.GetRolePermissionReq) (resp *types.Response, err error) {
	permissions, err := l.svcCtx.PermissionRpc.GetRolePermissions(l.ctx, &pb.RoleIdReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return &types.Response{Code: errno.ErrRPCFailed}, err
	}

	return &types.Response{Data: permissions}, nil
}
