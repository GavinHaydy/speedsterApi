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
	result, err := l.svcCtx.SysRoleModel.Insert(l.ctx, &model.Role{
		Name:        in.Name,
		Code:        in.Code,
		Description: utils.ToNullString(&in.Description),
	})
	if err != nil {
		return nil, errorx.New(errno.ErrInsertFailed)
	}

	roleId, err := result.LastInsertId()
	if err != nil {
		return nil, errorx.New(errno.ErrInsertFailed)
	}

	return &pb.CreateRoleResp{RoleId: roleId}, nil
}
