package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoGetLogic {
	return &UserInfoGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoGetLogic) UserInfoGet(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	one, err := l.svcCtx.SysUserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	return &pb.UserInfoResp{
		Id:       one.Id,
		Nickname: utils.NullString(one.Nickname),
	}, nil
}
