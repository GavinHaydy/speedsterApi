// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"
	"speedsterApi/common/response"

	"speedsterApi/app/role/internal/logic"
	"speedsterApi/app/role/internal/svc"
	"speedsterApi/app/role/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// DelRoleHandler 删除角色
func DelRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelRole
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelRoleLogic(r.Context(), svcCtx)
		resp, err := l.DelRole(&req)
		if err != nil {
			logx.Errorf("l.DelRole failed: %v", err)
			response.Error(w, r, resp.Code)
		} else {
			logx.Infof("l.DelRole success: %v", resp)
			response.Success(w, r)
		}
	}
}
