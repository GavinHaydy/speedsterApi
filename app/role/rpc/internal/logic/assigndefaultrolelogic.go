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
	_, err := l.svcCtx.SysUserRoleModel.Insert(l.ctx, &model.SysUserRole{
		UserId: in.UserId,
		RoleId: 1,
	})
	if err != nil {
		return nil, errorx.New(errno.ErrInsertFailed)
	}
	return &pb.AssignDefaultRoleResp{Code: errno.Ok}, nil
}
