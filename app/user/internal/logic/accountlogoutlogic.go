// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 退出登录
func NewAccountLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLogoutLogic {
	return &AccountLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLogoutLogic) AccountLogout() (resp *types.Rsp, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("user_id")

	logx.Infof("AccountLogoutLogic, Authorization:%v", value)
	return &types.Rsp{Data: value}, nil
}
