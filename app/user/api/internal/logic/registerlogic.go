// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/app/user/api/internal/types"
	"speedsterApi/app/user/user/user"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils"

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
	//_, err = l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, req.Username)
	//if err == nil {
	//	return &types.Response{Code: errno.ErrYetAccountRegister}, nil
	//}
	//
	//_, err = l.svcCtx.SysUserModel.FindOneByPhone(l.ctx, sql.NullString{
	//	String: req.Phone,
	//	Valid:  true,
	//})
	//if err == nil {
	//	return &types.Response{Code: errno.ErrYetPhoneRegister}, nil
	//}
	//
	//_, err = l.svcCtx.SysUserModel.FindOneByEmail(l.ctx, sql.NullString{
	//	String: req.Email,
	//	Valid:  true,
	//})
	//if err == nil {
	//	return &types.Response{Code: errno.ErrYetEmailRegister}, nil
	//}
	//
	//pw := utils.AesEncrypt(req.Password, l.svcCtx.Config.AesSecretKey)
	//_, err = l.svcCtx.SysUserModel.Insert(l.ctx, &model.SysUser{
	//	Id:       utils.GetUUID(),
	//	Username: req.Username,
	//	Password: pw,
	//	Phone: sql.NullString{
	//		String: req.Phone,
	//		Valid:  true,
	//	},
	//	Email: sql.NullString{
	//		String: req.Email,
	//		Valid:  true,
	//	},
	//	Nickname: sql.NullString{
	//		String: req.Username,
	//		Valid:  true,
	//	},
	//})
	//if err != nil {
	//	return &types.Response{
	//		Code: errno.ErrRegisterFailed,
	//	}, nil
	//}
	//
	//return &types.Response{}, nil
	register, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    utils.EmptyToNil(req.Phone),
		Email:    utils.EmptyToNil(req.Email),
	})
	logx.Infof("Register %+v", register)
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}
	return &types.Response{Data: register}, nil
}
