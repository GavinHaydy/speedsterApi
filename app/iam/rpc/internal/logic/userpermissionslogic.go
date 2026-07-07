package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermissionsLogic {
	return &UserPermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPermissionsLogic) UserPermissions(in *pb.UserPermissionsReq) (*pb.PermissionTreeResp, error) {
	logx.Infof("UserPermissions - %+v", in.UserId)
	roleId, err := l.svcCtx.SysUserRoleModel.FindRoleByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errorx.New(errno.ErrUserNotRole)
	}

	permissionIds, err := l.svcCtx.SysRolePermissionModel.FindByRoleId(l.ctx, roleId)
	if err != nil {
		return nil, errorx.New(errno.ErrRolePermissionEmpty)
	}

	logx.Infof("SysPermissionModel=%#v", l.svcCtx.SysPermissionModel)
	permissionList, err := l.svcCtx.SysPermissionModel.SelectTreeById(l.ctx, permissionIds)
	if err != nil {
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	userInfo, err := l.svcCtx.SysUserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}
	return &pb.PermissionTreeResp{
		List:    BuildPermissionTree(permissionList),
		IsAdmin: userInfo.IsSuper,
	}, nil
}
