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

// CreatePermissionHandler 新增权限
func CreatePermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePermissionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewCreatePermissionLogic(r.Context(), svcCtx)
		resp, err := l.CreatePermission(&req)
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.Success(w, r)
		}
	}
}
