// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"gateway/internal/logic"
	"gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DocPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDocPageLogic(r.Context(), svcCtx)
		resp, err := l.DocPage()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
