// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"speedsterApi/common/utils"
	"user/model"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRegisterLogic 注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, req.Username)
	if err == nil {
		return &types.Response{Msg: "username"}, nil // 没报错说明找到了，即用户名已存在
	}

	_, err = l.svcCtx.SysUserModel.FindOneByPhone(l.ctx, sql.NullString{
		String: req.Phone,
		Valid:  true,
	})
	if err == nil {
		return &types.Response{Msg: "phone"}, nil // 没报错说明找到了，即用户名已存在
	}

	_, err = l.svcCtx.SysUserModel.FindOneByEmail(l.ctx, sql.NullString{
		String: req.Email,
		Valid:  true,
	})
	if err == nil {
		return &types.Response{Msg: "email"}, nil // 没报错说明找到了，即用户名已存在
	}

	pw := utils.AesEncrypt(req.Password, l.svcCtx.Config.AesSecretKey)
	_, err = l.svcCtx.SysUserModel.Insert(l.ctx, &model.SysUser{
		Id:       utils.GetUUID(),
		Username: req.Username,
		Password: pw,
		Phone: sql.NullString{
			String: req.Phone,
			Valid:  true,
		},
		Email: sql.NullString{
			String: req.Email,
			Valid:  true,
		},
		Nickname: sql.NullString{
			String: req.Username,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.Response{Msg: "ok"}, nil
}
