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

// RefreshHandler 刷新token
func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh(&req)
		if err != nil {
			logx.Errorf("RefreshToken error: %v,---result%v", err, resp)
			response.ErrorWithCode(w, r, resp.Code)
		} else {
			logx.Infof("RefreshToken result: %v", resp)
			response.SuccessWithData(w, r, resp.Data)
		}
	}
}
