// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"user/internal/types"

	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLoginLogic {
	return &AccountLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLoginLogic) AccountLogin(req *types.LoginReq) (*types.Response, error) {
	// todo: add your logic here and delete this line
	//logx.Debug(l.ctx)
	//return nil
	return &types.Response{
		Msg: fmt.Sprintf("account:%s,password:%s", req.Account, req.Password),
	}, nil

}
