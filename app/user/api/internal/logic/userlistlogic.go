// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户列表
func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	//total, list, err := l.svcCtx.SysUserModel.SelectUserList(l.ctx, req)
	//if err != nil {
	//	return &types.Response{Code: errno.ErrSelectDbFailed}, err
	//}
	//result := map[string]interface{}{
	//	"list":  list,
	//	"total": total,
	//}
	//return &types.Response{Data: result}, nil
	return &types.Response{}, nil
}
