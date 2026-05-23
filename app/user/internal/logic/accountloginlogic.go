// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"
	"speedsterApi/common/errno"
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

	_, err := l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return &types.Response{Code: errno.ErrAccountNotFound}, err
	}

	pw := utils.AesEncrypt(req.Password, l.svcCtx.Config.AesSecretKey)

	userInfo, err := l.svcCtx.SysUserModel.FindByAccountAndPW(l.ctx, req.Username, pw)
	if err != nil {

		return &types.Response{Code: errno.ErrPasswordFailed}, err
	}

	var resultToken string
	var resultTime time.Time

	sysToken, err := rdb.Get(fmt.Sprintf("%s%v", l.svcCtx.Config.JWT.Prefix, userInfo.Id))

	if err != nil || sysToken == "" {
		logx.Info("token不存在:开始生成token")
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
			return &types.Response{Code: errno.ErrRoleNotExists}, nil
		}
		logx.Infof("role:%v", role)
		token, t, err := utils.GenerateToken(userInfo.Id, role, l.svcCtx.Config.JWT.Issuer, l.svcCtx.Config.JWT.Secret)
		if err != nil {
			logx.Errorf("GenerateToken: %v", err)
			return &types.Response{Code: errno.ErrGenTokenFailed}, nil
		}
		resultToken = token
		resultTime = t
	} else {
		logx.Info("token存在:开始刷新token")
		newToken, t, err := utils.RefreshToken(sysToken, l.svcCtx.Config.JWT.Issuer, l.svcCtx.Config.JWT.Secret)
		if err != nil {
			logx.Errorf("RefreshToken:%v", err)
			return &types.Response{Code: errno.ErrGenTokenFailed}, err
		}
		resultToken = newToken
		resultTime = t

	}

	exp := utils.DateTime(resultTime)

	err = rdb.Setex(fmt.Sprintf("%s%v", l.svcCtx.Config.JWT.Prefix, userInfo.Id), resultToken, l.svcCtx.Config.JWT.Expire)
	if err != nil {
		logx.Errorf("Setex: %v", err)
		return &types.Response{Code: errno.ErrRedisFailed}, err
	}

	result := map[string]interface{}{
		"token": resultToken,
		"time":  exp,
	}
	return &types.Response{
		Data: result,
	}, nil

}
