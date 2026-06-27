package logic

import (
	"context"
	"speedsterApi/app/iam/model"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserRoleLogic {
	return &AddUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserRoleLogic) AddUserRole(in *pb.UserRole) (*pb.UserRoleResp, error) {
	_, err := l.svcCtx.SysUserRoleModel.Insert(l.ctx, &model.SysUserRole{
		UserId: in.UserId,
		RoleId: in.RoleId,
	})

	if err != nil {
		return nil, errorx.New(errno.ErrInsertFailed)
	}

	return &pb.UserRoleResp{Code: errno.Ok}, nil
}
