// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/common/response"
	"user/internal/types"

	"user/internal/logic"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AccountLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		// 1. go-zero 自动将请求的 JSON 解析到 req 结构体中
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAccountLoginLogic(r.Context(), svcCtx)
		result, err := l.AccountLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.SuccessWithData(w, result.Data)
		}
	}
}
