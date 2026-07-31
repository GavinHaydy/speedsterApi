// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"
	"speedsterApi/app/iam/rpc/iam"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserInfoLogic 用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.EmptyReq) (resp *types.Response, err error) {
	userInfo, err := l.svcCtx.IamRpc.UserInfoGet(l.ctx, &iam.UserInfoReq{
		UserId: l.ctx.Value("user_id").(string),
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}

	return &types.Response{Data: userInfo}, nil
}
