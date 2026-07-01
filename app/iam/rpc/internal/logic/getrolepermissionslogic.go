package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionsLogic {
	return &GetRolePermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRolePermissionsLogic) GetRolePermissions(in *pb.RoleIdReq) (*pb.RolePermissionResp, error) {
	ids, err := l.svcCtx.SysRolePermissionModel.FindByRoleId(l.ctx, in.RoleId)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GetRolePermissions error:%+v", err)
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	return &pb.RolePermissionResp{
		PermissionIds: ids,
	}, nil
}
