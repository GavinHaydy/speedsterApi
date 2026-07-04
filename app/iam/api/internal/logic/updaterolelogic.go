// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils/ptr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateRoleLogic 修改角色
func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.Response, err error) {
	_, err = l.svcCtx.IamRpc.RoleUpdate(l.ctx, &pb.UpdateRoleReq{
		Id:          req.Id,
		Status:      ptr.Value(req.Status),
		Name:        ptr.Value(req.Name),
		Description: ptr.Value(req.Description),
		Code:        ptr.Value(req.Code),
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}

	return nil, nil
}
