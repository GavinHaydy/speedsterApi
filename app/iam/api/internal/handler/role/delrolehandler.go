// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package role

import (
	"net/http"
	"speedsterApi/app/iam/api/internal/logic/role"
	"speedsterApi/common/response"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// DelRoleHandler 删除角色
func DelRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewDelRoleLogic(r.Context(), svcCtx)
		resp, err := l.DelRole(&req)
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.Success(w, r)
		}
	}
}
