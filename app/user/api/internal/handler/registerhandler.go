// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/app/user/api/internal/logic"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/app/user/api/internal/types"
	"speedsterApi/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// RegisterHandler 注册
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.ErrorWithCode(w, r, resp.Code)
		} else {
			response.Success(w, r)
		}
	}
}
