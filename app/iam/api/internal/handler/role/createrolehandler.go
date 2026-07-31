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

// CreateRoleHandler 新建角色
func CreateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewCreateRoleLogic(r.Context(), svcCtx)
		resp, err := l.CreateRole(&req)
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.SuccessWithData(w, r, resp.Data)
		}
	}
}
