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

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *pb.UpdateRoleReq) (*pb.Empty, error) {
	role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	updateRole := &model.Role{
		Id:          role.Id,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		Status:      role.Status,
	}
	if in.Name != "" {
		updateRole.Name = in.Name
	}
	if in.Code != "" {
		updateRole.Code = in.Code
	}
	if in.Description != "" {
		updateRole.Description = utils.ToNullString(&in.Description)
	}
	if in.Status != 0 {
		updateRole.Status = in.Status
	}

	err = l.svcCtx.SysRoleModel.Update(l.ctx, updateRole)
	if err != nil {
		return nil, errorx.New(errno.ErrUpdateDataFailed)
	}

	//return &types.Response{Code: errno.Ok}, nil

	return &pb.Empty{}, nil
}
