// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/common/errno"

	"speedsterApi/app/user/internal/svc"
	"speedsterApi/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewStatusLogic 修改用户状态
func NewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatusLogic {
	return &StatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatusLogic) Status(req *types.StatusReq) (resp *types.Response, err error) {
	userInfo, err := l.svcCtx.SysUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return &types.Response{Code: errno.ErrRecordNotFound}, err
	}
	userInfo.Status = req.Status

	err = l.svcCtx.SysUserModel.Update(l.ctx, userInfo)
	if err != nil {
		return &types.Response{Code: errno.ErrUpdateDataFailed}, err
	}

	return nil, nil
}
