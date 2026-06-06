// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/app/user/internal/logic"
	"speedsterApi/app/user/internal/svc"
	"speedsterApi/app/user/internal/types"
	"speedsterApi/common/response"

	"github.com/zeromicro/go-zero/core/logx"
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
			logx.Errorf("AccountLoginLogic error: %v,---result%v", err, result)
			response.ErrorWithCode(w, r, result.Code)
		} else {
			logx.Infof("AccountLoginLogic result: %v", result)
			response.SuccessWithData(w, r, result.Data)
		}
	}
}
