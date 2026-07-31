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

type DelRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDelRoleLogic 删除角色
func NewDelRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRoleLogic {
	return &DelRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRoleLogic) DelRole(req *types.DelRoleReq) (resp *types.Response, err error) {
	_, err = l.svcCtx.IamRpc.RoleDelete(l.ctx, &pb.DelRoleReq{
		Id: req.Id,
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
