// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package permission

import (
	"context"
	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"
	"speedsterApi/app/iam/rpc/pb"
	"speedsterApi/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreatePermissionLogic 新增权限
func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePermissionLogic) CreatePermission(req *types.CreatePermissionReq) (resp *types.Response, err error) {
	_, err = l.svcCtx.IamRpc.PermissionCreate(l.ctx, &pb.CreatePermissionReq{
		ParentId: req.ParentId,
		Name:     req.Name,
		Code:     req.Code,
		Path:     req.Path,
		Method:   req.Method,
		Type:     req.Type,
		Icon:     req.Icon,
		Sort:     req.Sort,
		Status:   req.Status,
	})
	if err != nil {
		code, msg := errorx.Parse(err)
		return &types.Response{
			Code: code,
			Msg:  msg,
		}, err
	}

	return nil, nil
}
