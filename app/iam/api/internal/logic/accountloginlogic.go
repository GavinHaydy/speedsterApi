// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAccountLoginLogic 登录
func NewAccountLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLoginLogic {
	return &AccountLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLoginLogic) AccountLogin(req *types.LoginReq) (*types.Response, error) {
	resp, err := l.svcCtx.IamRpc.Login(
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
		Data: map[string]interface{}{
			"accessToken":  resp.AccessToken,
			"refreshToken": resp.RefreshToken,
			"time":         resp.Time,
		},
	}, nil

}
