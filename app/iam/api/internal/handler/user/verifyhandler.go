// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"net/http"
	"speedsterApi/common/response"

	"speedsterApi/app/iam/api/internal/logic/user"
	"speedsterApi/app/iam/api/internal/svc"
)

// VerifyHandler 临时验证
func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify()
		if err != nil {
			response.Error(w, r, resp.Code)
		} else {
			response.Success(w, r)
		}
	}
}
