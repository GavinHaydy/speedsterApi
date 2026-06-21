// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"speedsterApi/app/iam/api/internal/logic"
	"speedsterApi/app/iam/api/internal/svc"
)

// 退出登录
func AccountLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAccountLogoutLogic(r.Context(), svcCtx)
		resp, err := l.AccountLogout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
