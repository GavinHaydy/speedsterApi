package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *pb.DelRoleReq) (*pb.Empty, error) {
	count, err := l.svcCtx.SysRolePermissionModel.FindByRoleId(l.ctx, in.Id)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("del role error, %v", err)
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	if count > 0 {
		return nil, errorx.New(errno.ErrRoleNotDel)
	}

	err = l.svcCtx.SysRoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.New(errno.ErrDeleteFailed)
	}

	return &pb.Empty{}, nil
}
