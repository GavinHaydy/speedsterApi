// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"
	"fmt"
	"speedsterApi/common/errno"
	"speedsterApi/common/utils"
	"time"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRefreshLogic 刷新token
func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (resp *types.Response, err error) {
	userId, _, tokenType, err := utils.ParseToken(req.RefreshToken, l.svcCtx.Config.Auth.RefreshSecret)
	if err != nil {
		return &types.Response{Code: errno.ErrSignError}, err
	}
	if tokenType != "refresh" {
		return &types.Response{Code: errno.ErrTokenTypeFailed}, errors.New("ErrTokenTypeFailed")
	}
	key := fmt.Sprintf("%s%s", l.svcCtx.Config.Auth.Prefix, userId)

	rdsToken, err := l.svcCtx.Redis.Get(key)
	if err != nil || rdsToken != req.RefreshToken {
		logx.Errorf("redis get token error: %v", err)
		return &types.Response{Code: errno.ErrInvalidToken}, errors.New("InvalidToken")
	}

	var accessToken string
	var refreshToken string
	var resultTime time.Time

	logx.Info("开始生成token")
	var role string

	// language=PostgresSQL
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
		userId,
	)
	if err != nil {
		return &types.Response{Code: errno.ErrRoleNotExists}, err
	}

	token, t, err := utils.GenAccessToken(userId, role, l.svcCtx.Config.Auth.Issuer, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Errorf("GenAccessToken: %v", err)
		return &types.Response{Code: errno.ErrGenTokenFailed}, err
	}
	accessToken = token
	resultTime = t

	longToken, _, err := utils.GenRefreshToken(userId, role, l.svcCtx.Config.Auth.Issuer, l.svcCtx.Config.Auth.RefreshSecret, l.svcCtx.Config.Auth.RefreshExpire)
	if err != nil {
		logx.Errorf("GenRefreshToken:%v", err)
		return &types.Response{Code: errno.ErrGenTokenFailed}, err
	}
	refreshToken = longToken

	//}

	exp := utils.DateTime(resultTime)

	err = l.svcCtx.Redis.Setex(fmt.Sprintf("%s%v", l.svcCtx.Config.Auth.Prefix, userId), refreshToken, l.svcCtx.Config.Auth.RefreshExpire)
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
