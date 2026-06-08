package logic

import (
	"context"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"speedsterApi/app/user/user/internal/svc"
	"speedsterApi/app/user/user/pb/pb"

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
		logx.Errorf("UserListLogic.err:%v", err)
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}
	logx.Infof("UserListLogic.total:%v,list:%v", total, list)

	return &pb.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
