// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/common/response"

	"speedsterApi/app/role/internal/logic"
	"speedsterApi/app/role/internal/svc"
	"speedsterApi/app/role/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateRoleHandler 新建角色
func CreateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewRole
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateRoleLogic(r.Context(), svcCtx)
		resp, err := l.CreateRole(&req)
		if err != nil {
			logx.Errorf("CreateRole err: %v", err)
			response.ErrorWithCode(w, r, resp.Code)
		} else {
			logx.Infof("CreateRole resp: %v", resp)
			response.Success(w, r)
		}
	}
}
