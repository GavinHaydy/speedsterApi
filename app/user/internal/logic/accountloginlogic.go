// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"speedsterApi/app/user/internal/svc"
	"speedsterApi/app/user/internal/types"
	"speedsterApi/common/errno"
	"time"

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

	_, err := l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return &types.Response{Code: errno.ErrAccountNotFound}, err
	}

	pw := utils.AesEncrypt(req.Password, l.svcCtx.Config.AesSecretKey)

	userInfo, err := l.svcCtx.SysUserModel.FindByAccountAndPW(l.ctx, req.Username, pw)
	if err != nil {

		return &types.Response{Code: errno.ErrPasswordFailed}, err
	}

	var accessToken string
	var refreshToken string
	var resultTime time.Time

	logx.Info("开始生成token")
	var role string

	// language=PostgreSQL
	sql := `
				select r.code
				from sys_user_role ur
				join role r on ur.role_id = r.id
				where ur.user_id = $1
				limit 1
				`

	err = l.svcCtx.DB.QueryRowCtx(
		l.ctx,
		&role,
		sql,
		userInfo.Id,
	)
	if err != nil {
		return &types.Response{Code: errno.ErrRoleNotExists}, err
	}

	token, t, err := utils.GenAccessToken(userInfo.Id, role, l.svcCtx.Config.Auth.Issuer, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Errorf("GenAccessToken: %v", err)
		return &types.Response{Code: errno.ErrGenTokenFailed}, nil
	}
	accessToken = token
	resultTime = t

	longToken, _, err := utils.GenRefreshToken(userInfo.Id, role, l.svcCtx.Config.Auth.Issuer, l.svcCtx.Config.Auth.RefreshSecret, l.svcCtx.Config.Auth.RefreshExpire)
	if err != nil {
		logx.Errorf("GenRefreshToken:%v", err)
		return &types.Response{Code: errno.ErrGenTokenFailed}, err
	}
	refreshToken = longToken

	//}

	exp := utils.DateTime(resultTime)

	err = rdb.Setex(fmt.Sprintf("%s%v", l.svcCtx.Config.Auth.Prefix, userInfo.Id), refreshToken, l.svcCtx.Config.Auth.RefreshExpire)
	if err != nil {
		logx.Errorf("Setex: %v", err)
		return &types.Response{Code: errno.ErrRedisFailed}, err
	}

	result := map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"time":         exp,
	}
	return &types.Response{
		Data: result,
	}, nil

}
