package logic

import (
	"context"
	"database/sql"
	"errors"
	"speedsterApi/app/iam/model"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

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

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterRsp, error) {
	roleLogic := NewAssignDefaultRoleLogic(l.ctx, l.svcCtx)

	_, err := l.svcCtx.SysUserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, errorx.New(errno.ErrYetAccountRegister)
	}

	pw := utils.AesEncrypt(in.Password, l.svcCtx.Config.CacheAuth.AccessSecret)

	userInfo := &model.SysUser{
		Id:       utils.GetUUID(),
		Username: in.Username,
		Password: pw,
		Status:   1,
	}

	logx.Infof("Register user: %+v", userInfo.Phone)

	if in.Phone != nil {
		_, err = l.svcCtx.SysUserModel.FindOneByPhone(l.ctx, utils.ToNullString(in.Phone))

		if err == nil {
			return nil, errorx.New(errno.ErrYetPhoneRegister)
		}

		if !errors.Is(err, model.ErrNotFound) {
			return nil, errorx.New(errno.ErrPgsqlFailed)
		}
		userInfo.Phone = utils.ToNullString(in.Phone)
	}

	if in.Email != nil {
		_, err = l.svcCtx.SysUserModel.FindOneByEmail(l.ctx, utils.ToNullString(in.Email))
		if err == nil {
			return nil, errorx.New(errno.ErrYetEmailRegister)
		}

		if !errors.Is(err, model.ErrNotFound) {
			return nil, errorx.New(errno.ErrPgsqlFailed)
		}
		userInfo.Email = utils.ToNullString(in.Email)
	}

	userInfo.Nickname = sql.NullString{
		String: in.Username,
		Valid:  true,
	}

	_, err = l.svcCtx.SysUserModel.Insert(l.ctx, userInfo)
	if err != nil {
		return nil, errorx.New(errno.ErrRegisterFailed)
	}

	logx.Infof("Register user: %+v", userInfo.Id)

	_, err = roleLogic.AssignDefaultRole(&pb.AssignDefaultRoleReq{UserId: userInfo.Id})
	if err != nil {
		return nil, errorx.New(errno.ErrUserNotRole)
	}

	return &pb.RegisterRsp{UserId: userInfo.Id}, nil
}
