// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"user/internal/svc"
	"user/internal/types"

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
	// todo: add your logic here and delete this line
	value := l.ctx.Value("user_id")
	_, err = l.svcCtx.Redis.Del(fmt.Sprintf("%s%v", l.svcCtx.Config.JWT.Prefix, value))
	if err != nil {
		return nil, err
	}

	return &types.Response{Msg: "logout success"}, nil
}
