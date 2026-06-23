package logic

import (
	"context"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignDefaultRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignDefaultRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignDefaultRoleLogic {
	return &AssignDefaultRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssignDefaultRoleLogic) AssignDefaultRole(in *pb.AssignDefaultRoleReq) (*pb.AssignDefaultRoleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AssignDefaultRoleResp{}, nil
}
