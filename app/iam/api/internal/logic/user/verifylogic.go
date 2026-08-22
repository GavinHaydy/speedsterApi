// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewVerifyLogic 临时验证
func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify() (resp *types.Response, err error) {

	return nil, nil
}
