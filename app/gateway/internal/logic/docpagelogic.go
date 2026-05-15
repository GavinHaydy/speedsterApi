// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DocPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DocPageLogic {
	return &DocPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DocPageLogic) DocPage() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
