package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpUserStatusLogic {
	return &UpUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpUserStatusLogic) UpUserStatus(in *pb.UpUserStatusReq) (*pb.UpUserStatusResp, error) {
	userInfo, err := l.svcCtx.SysUserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errorx.New(errno.ErrRecordNotFound)
	}
	userInfo.Status = in.Status

	err = l.svcCtx.SysUserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, errorx.New(errno.ErrUpdateDataFailed)
	}

	return &pb.UpUserStatusResp{Code: 1}, nil
}
