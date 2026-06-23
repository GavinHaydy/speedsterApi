package logic

import (
	"context"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionTreeLogic {
	return &PermissionTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionTreeLogic) PermissionTree(in *pb.PermissionTreeReq) (*pb.PermissionTreeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PermissionTreeResp{}, nil
}
