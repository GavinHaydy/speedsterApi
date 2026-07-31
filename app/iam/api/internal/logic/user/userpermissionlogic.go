// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserPermissionLogic 用户权限列表
func NewUserPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermissionLogic {
	return &UserPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPermissionLogic) UserPermission() (resp *types.Response, err error) {
	result, err := l.svcCtx.IamRpc.UserPermissions(l.ctx, &pb.UserPermissionsReq{
		UserId: l.ctx.Value("user_id").(string),
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}
	return &types.Response{Data: result}, nil
}
