// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"speedsterApi/common/errno"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAccountLogoutLogic 退出登录
func NewAccountLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLogoutLogic {
	return &AccountLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLogoutLogic) AccountLogout() (resp *types.Response, err error) {
	value := l.ctx.Value("user_id")
	_, err = l.svcCtx.Redis.Del(fmt.Sprintf("%s%v", l.svcCtx.Config.Auth.Prefix, value))
	if err != nil {
		return &types.Response{Code: errno.ErrRedisFailed}, err
	}
	return &types.Response{}, nil
}
