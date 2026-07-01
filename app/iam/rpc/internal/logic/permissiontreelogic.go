package logic

import (
	"context"
	"speedsterApi/app/iam/model"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"
	"speedsterApi/common/utils"

	"speedsterApi/app/iam/rpc/internal/svc"
	"speedsterApi/app/iam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionTreeLogic {
	return &PermissionTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func BuildPermissionTree(
	list []*model.SysPermission,
) []*pb.PermissionTreeItem {

	nodeMap := make(map[int64]*pb.PermissionTreeItem)

	for _, item := range list {
		nodeMap[item.Id] = &pb.PermissionTreeItem{
			Id:       item.Id,
			Name:     item.Name,
			Code:     item.Code,
			Path:     utils.NullString(item.Path),
			Method:   utils.NullString(item.Method),
			Type:     item.Type,
			Icon:     item.Icon.String,
			Sort:     item.Sort,
			Status:   item.Status,
			ParentId: item.ParentId,
			Children: []*pb.PermissionTreeItem{},
		}
	}

	var roots []*pb.PermissionTreeItem

	for _, node := range nodeMap {

		if node.ParentId == 0 {
			roots = append(roots, node)
			continue
		}

		parent, ok := nodeMap[node.ParentId]
		if !ok {
			continue
		}

		parent.Children = append(
			parent.Children,
			node,
		)
	}

	return roots
}

func (l *PermissionTreeLogic) PermissionTree(in *pb.PermissionTreeReq) (*pb.PermissionTreeResp, error) {
	list, err := l.svcCtx.SysPermissionModel.SelectTree(
		l.ctx,
		in,
	)

	if err != nil {
		logx.Errorf("l.svcCtx.SysPermission.SelectTree err:%v", err)
		return nil, errorx.New(errno.ErrSelectDbFailed)
	}

	return &pb.PermissionTreeResp{
		List: BuildPermissionTree(list),
	}, nil
}
