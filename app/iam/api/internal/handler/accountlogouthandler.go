// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/common/response"

	"speedsterApi/app/iam/api/internal/logic"
	"speedsterApi/app/iam/api/internal/svc"
)

// AccountLogoutHandler 退出登录
func AccountLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAccountLogoutLogic(r.Context(), svcCtx)
		resp, err := l.AccountLogout()
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.Success(w, r)
		}
	}
}
