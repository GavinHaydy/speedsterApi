// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/app/role/api/internal/logic"
	"speedsterApi/app/role/api/internal/svc"
	"speedsterApi/app/role/api/internal/types"
	"speedsterApi/common/response"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// UpdateRoleHandler 修改角色
func UpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateRoleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRole(&req)
		if err != nil {
			logx.Errorf("%+v", resp)
			response.ErrorWithCode(w, r, resp.Code)
		} else {
			logx.Infof("%+v", resp)
			response.Success(w, r)
		}
	}
}
