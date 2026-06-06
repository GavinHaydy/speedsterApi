// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/app/role/api/internal/logic"
	"speedsterApi/app/role/api/internal/svc"
	"speedsterApi/app/role/api/internal/types"
	"speedsterApi/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// RoleListHandler 角色列表
func RoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRoleListLogic(r.Context(), svcCtx)
		resp, err := l.RoleList(&req)
		if err != nil {
			response.ErrorWithCode(w, r, resp.Code)
		} else {
			response.SuccessWithData(w, r, resp.Data)
		}
	}
}
