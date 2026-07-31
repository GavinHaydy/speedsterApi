// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"net/http"
	"speedsterApi/app/iam/api/internal/logic/user"
	"speedsterApi/common/response"

	"speedsterApi/app/iam/api/internal/svc"
	"speedsterApi/app/iam/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// UserPermissionHandler 用户权限列表
func UserPermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserPermissionLogic(r.Context(), svcCtx)
		resp, err := l.UserPermission()
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.SuccessWithData(w, r, resp.Data)
		}
	}
}
