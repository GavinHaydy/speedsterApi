package logic

import (
	"context"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleCreateLogic {
	return &RoleCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleCreateLogic) RoleCreate(in *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateRoleResp{}, nil
}
