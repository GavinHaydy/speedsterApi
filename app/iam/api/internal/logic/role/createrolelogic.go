// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package role

import (
	"context"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreateRoleLogic 新建角色
func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.NewRoleReq) (resp *types.Response, err error) {
	roleId, err := l.svcCtx.IamRpc.RoleCreate(l.ctx, &pb.CreateRoleReq{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}

	return &types.Response{Data: roleId}, nil
}
