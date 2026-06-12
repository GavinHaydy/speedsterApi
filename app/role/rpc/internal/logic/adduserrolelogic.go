package logic

import (
	"context"
	"speedsterApi/app/role/model"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/role/rpc/internal/svc"
	"speedsterApi/app/role/rpc/pb"

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
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.SysUserRoleModel.Insert(l.ctx, &model.SysUserRole{
		UserId: in.UserId,
		RoleId: in.RoleId,
	})
	if err != nil {
		return nil, errorx.New(errno.ErrRegisterFailed)
	}

	return &pb.UserRoleResp{Code: errno.Ok}, nil
}
