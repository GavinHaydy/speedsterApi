// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package permission

import (
	"net/http"
	"speedsterApi/app/iam/api/internal/logic/permission"
	"speedsterApi/common/response"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetRolePermissionsHandler 角色权限
func GetRolePermissionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRolePermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewGetRolePermissionsLogic(r.Context(), svcCtx)
		resp, err := l.GetRolePermissions(&req)
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.SuccessWithData(w, r, resp.Data)
		}
	}
}
