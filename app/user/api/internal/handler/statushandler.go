// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/app/user/api/internal/logic"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/app/user/api/internal/types"
	"speedsterApi/common/response"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// StatusHandler 修改用户状态
func StatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStatusLogic(r.Context(), svcCtx)
		resp, err := l.Status(&req)
		if err != nil {
			logx.Errorf("StatusLogic err: %+v", err)
			response.ErrorWithCode(w, r, resp.Code)
		} else {
			response.Success(w, r)
		}
	}
}
