package logic

import (
	"context"
	"speedsterApi/app/iam/model"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionCreateLogic {
	return &PermissionCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionCreateLogic) PermissionCreate(in *pb.CreatePermissionReq) (*pb.Empty, error) {
	//l.svcCtx.SysPermissionModel.Insert(l.ctx,)
	newPermission := &model.SysPermission{
		ParentId: in.ParentId,
		Name:     in.Name,
		Code:     in.Code,
		Type:     in.Type,
		Sort:     in.Sort,
	}
	if in.Path != nil {
		newPermission.Path = utils.ToNullString(in.Path)
	}
	if in.Method != nil {
		newPermission.Method = utils.ToNullString(in.Method)
	}
	if in.Icon != nil {
		newPermission.Icon = utils.ToNullString(in.Icon)
	}
	if in.Status != nil {
		newPermission.Status = *in.Status
	}
	_, err := l.svcCtx.SysPermissionModel.Insert(l.ctx, newPermission)
	if err != nil {
		return nil, errorx.New(errno.ErrInsertFailed)
	}
	return &pb.Empty{}, nil
}
