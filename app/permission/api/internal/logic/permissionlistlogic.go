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

type PermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewPermissionListLogic 权限列表
func NewPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListLogic {
	return &PermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionListLogic) PermissionList(req *types.PermisssionListReq) (resp *types.Response, err error) {
	tree, err := l.svcCtx.PermissionRpc.PermissionTree(l.ctx, &pb.PermissionTreeReq{
		Name: req.Name,
	})
	if err != nil {
		return &types.Response{Code: errno.ErrRPCFailed}, err
	}

	return &types.Response{Data: tree}, nil
}
