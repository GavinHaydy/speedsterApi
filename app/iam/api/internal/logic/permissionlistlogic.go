// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

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
	tree, err := l.svcCtx.IamRpc.PermissionTree(l.ctx, &pb.PermissionTreeReq{
		Name: req.Name,
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{Code: code, Msg: msg}, err
	}

	return &types.Response{Data: tree}, nil
}
