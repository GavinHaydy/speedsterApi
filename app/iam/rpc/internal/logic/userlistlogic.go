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
		logx.Errorf("UserList,error:%+v", err)
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}
	logx.Infof("UserList,total:%+v,list:%+v", total, list)

	return &pb.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
