// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/app/user/api/internal/types"
	"speedsterApi/app/user/user/pb/pb"
	"speedsterApi/common/errorx"

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

	resp, err := l.svcCtx.UserRpc.Login(
		l.ctx,
		&pb.LoginReq{Username: req.Username, Password: req.Password},
	)

	if err != nil {
		logx.Errorf("---------------error: %+v", err.Error())
		code, msg := errorx.Parse(err)

		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}

	return &types.Response{
		Data: resp,
	}, nil

}
