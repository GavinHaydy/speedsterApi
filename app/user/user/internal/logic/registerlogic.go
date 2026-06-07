package logic

import (
	"context"
	"database/sql"
	"speedsterApi/app/user/model"
	"speedsterApi/app/user/user/internal/svc"
	"speedsterApi/app/user/user/user"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRsp, error) {
	_, err := l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return nil, errorx.New(errno.ErrYetAccountRegister)
	}

	_, err = l.svcCtx.SysUserModel.FindOneByPhone(l.ctx, sql.NullString{
		String: in.Phone,
		Valid:  true,
	})
	if err == nil {
		return nil, errorx.New(errno.ErrYetPhoneRegister)
	}

	_, err = l.svcCtx.SysUserModel.FindOneByEmail(l.ctx, sql.NullString{
		String: in.Email,
		Valid:  true,
	})
	if err == nil {
		return nil, errorx.New(errno.ErrYetEmailRegister)
	}

	pw := utils.AesEncrypt(in.Password, l.svcCtx.Config.CacheAuth.AccessSecret)
	userInfo, err := l.svcCtx.SysUserModel.Insert(l.ctx, &model.SysUser{
		Id:       utils.GetUUID(),
		Username: in.Username,
		Password: pw,
		Phone: sql.NullString{
			String: in.Phone,
			Valid:  true,
		},
		Email: sql.NullString{
			String: in.Email,
			Valid:  true,
		},
		Nickname: sql.NullString{
			String: in.Username,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, errorx.New(errno.ErrRegisterFailed)
	}
	id, err := userInfo.LastInsertId()
	if err != nil {
		return nil, err
	}
	logx.Infof("register user:%+v", id)

	return &user.RegisterRsp{}, nil
	//return &user.RegisterRsp{}, nil
}
