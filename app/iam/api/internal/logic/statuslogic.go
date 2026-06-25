// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"

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
	_, err = l.svcCtx.IamRpc.UpUserStatus(l.ctx, &pb.UpUserStatusReq{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		code, msg := errorx.Parse(err)

		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}

	return &types.Response{}, nil
}
