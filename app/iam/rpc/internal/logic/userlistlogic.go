package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *pb.UserListReq) (*pb.UserListResp, error) {
	total, list, err := l.svcCtx.SysUserModel.SelectUserList(l.ctx, in)
	if err != nil {
		logx.Errorw("UserList", logx.Field("error", err.Error()))
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}
	logx.Infow("UserList", logx.Field("list", list), logx.Field("total", total))

	return &pb.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
