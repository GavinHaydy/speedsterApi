package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *pb.RoleListReq) (*pb.RoleListResp, error) {
	list, err := l.svcCtx.SysRoleModel.SelectRoleList(l.ctx, in)
	if err != nil {
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	return list, nil
}
