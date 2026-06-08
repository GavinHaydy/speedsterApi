// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/app/user/api/internal/types"
	"speedsterApi/app/user/user/pb/pb"
	"speedsterApi/common/errno"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserListLogic 用户列表
func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.Response, err error) {
	data, err := l.svcCtx.UserRpc.UserList(l.ctx, &pb.UserListReq{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Status:   req.Status,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		logx.Errorf("------------%v", err.Error())
		return &types.Response{Code: errno.ErrRPCFailed}, err
	}
	logx.Infof("UserListReq: %+v", data)
	return &types.Response{Data: map[string]interface{}{
		"total": data.Total,
		"list":  data.List,
	}}, nil
}
