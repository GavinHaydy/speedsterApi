// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/common/errno"
	"speedsterApi/common/response"

	"user/internal/logic"
	"user/internal/svc"
)

// AccountLogoutHandler 退出登录
func AccountLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAccountLogoutLogic(r.Context(), svcCtx)
		_, err := l.AccountLogout()
		if err != nil {
			response.ErrorWithCode(w, r, errno.ErrServer)
		} else {
			response.Success(w, r)
		}
	}
}
