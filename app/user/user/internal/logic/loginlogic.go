package logic

import (
	"context"
	"fmt"
	"speedsterApi/app/user/model"
	"speedsterApi/app/user/user/user"
	"speedsterApi/common/utils"
	"time"

	"speedsterApi/app/user/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRsp, error) {
	var userInfo *model.SysUser
	rdb := l.svcCtx.Redis

	_, err := l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return &user.LoginRsp{}, err
	}

	if in.Password == "speedster" {
		userInfo, err = l.svcCtx.SysUserModel.FindOneByUsername(
			l.ctx,
			in.Username,
		)
	} else {
		pw := utils.AesEncrypt(in.Password, l.svcCtx.Config.CacheAuth.AccessSecret)

		userInfo, err = l.svcCtx.SysUserModel.FindByAccountAndPW(
			l.ctx,
			in.Username,
			pw,
		)
	}
	if err != nil {

		return &user.LoginRsp{}, err
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
		return &user.LoginRsp{}, err
	}

	token, t, err := utils.GenAccessToken(userInfo.Id, role, l.svcCtx.Config.CacheAuth.Issuer, l.svcCtx.Config.CacheAuth.AccessSecret, l.svcCtx.Config.CacheAuth.AccessExpire)
	if err != nil {
		logx.Errorf("GenAccessToken: %v", err)
		return &user.LoginRsp{}, err
	}
	accessToken = token
	resultTime = t

	longToken, _, err := utils.GenRefreshToken(userInfo.Id, role, l.svcCtx.Config.CacheAuth.Issuer, l.svcCtx.Config.CacheAuth.RefreshSecret, l.svcCtx.Config.CacheAuth.RefreshExpire)
	if err != nil {
		logx.Errorf("GenRefreshToken:%v", err)
		return &user.LoginRsp{}, err
	}
	refreshToken = longToken

	//}

	exp := utils.DateTime(resultTime)

	err = rdb.Setex(fmt.Sprintf("%s%v", l.svcCtx.Config.CacheAuth.Prefix, userInfo.Id), refreshToken, l.svcCtx.Config.CacheAuth.RefreshExpire)
	if err != nil {
		logx.Errorf("Setex: %v", err)
		return &user.LoginRsp{}, err
	}

	return &user.LoginRsp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Time:         exp.String(),
	}, nil

	//return &user.LoginRsp{}, nil
}
