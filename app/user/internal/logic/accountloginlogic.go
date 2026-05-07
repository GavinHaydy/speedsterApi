// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"time"
	"user/internal/types"

	"user/internal/svc"

	"speedsterApi/common/utils"

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
	rdb := l.svcCtx.Redis

	pw := utils.AesEncrypt(req.Password, l.svcCtx.Config.AesSecretKey)

	userInfo, err := l.svcCtx.SysUserModel.FindByAccountAndPW(l.ctx, req.Account, pw)
	if err != nil {
		return &types.Response{Msg: err.Error()}, nil
	}

	var resultToken string
	var resultTime time.Time

	sysToken, err := rdb.Get(fmt.Sprintf("token_%v", userInfo.Phone.String))

	logx.Infof("sysToken:%v, err: %v", sysToken, err)
	if err != nil || sysToken == "" {
		logx.Info("token不存在:开始生成token")
		logx.Infof("===================================%v", sysToken)
		token, t, err := utils.GenerateToken(userInfo.Id, l.svcCtx.Config.JWT.Issuer, l.svcCtx.Config.JWT.Secret)
		if err != nil {
			logx.Debugf("===%v", err)
			return &types.Response{Msg: err.Error()}, err
		}
		resultToken = token
		resultTime = t
	} else {
		logx.Infof("------------------------%v", sysToken)
		logx.Info("token存在:开始刷新token")
		newToken, t, err := utils.RefreshToken(sysToken, l.svcCtx.Config.JWT.Issuer, l.svcCtx.Config.JWT.Secret)
		if err != nil {
			logx.Infof("===%v", err)
			return &types.Response{Msg: err.Error()}, err
		}
		resultToken = newToken
		resultTime = t

	}

	exp := utils.DateTime(resultTime)

	err = rdb.Setex(fmt.Sprintf("token_%v", userInfo.Phone.String), resultToken, l.svcCtx.Config.JWT.Expire)
	if err != nil {
		return &types.Response{Msg: err.Error()}, err
	}

	result := map[string]interface{}{
		"token": resultToken,
		"time":  exp,
	}
	return &types.Response{
		Code: 0,
		Msg:  "",
		Data: result,
	}, nil

}
