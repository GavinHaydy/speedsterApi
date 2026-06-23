package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.UserRoleResp{}, nil
}
