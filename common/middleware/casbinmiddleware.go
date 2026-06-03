package middleware

import (
	"errors"
	"net/http"

	mycasbin "speedsterApi/common/casbin"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CasbinMiddleware struct{}

func NewCasbinMiddleware() *CasbinMiddleware {
	return &CasbinMiddleware{}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("----------CasbinMiddleware")

		role, ok := r.Context().Value("role").(string)
		if !ok {
			httpx.Error(w, errors.New("no role"))
			return
		}
		if role == "admin" {
			next(w, r)
			return
		}

		pass, err := mycasbin.Enforcer.Enforce(
			role,
			r.URL.Path,
			r.Method,
		)
		logx.Infof(
			"role=%v path=%s method=%s pass=%v err=%v",
			role,
			r.URL.Path,
			r.Method,
			pass,
			err,
		)

		if err != nil {
			logx.Errorf("CasbinMiddleware err: %v", err)
			httpx.Error(w, err)
			return
		}

		if !pass {
			httpx.Error(w, errors.New("permission denied"))
			return
		}

		next(w, r)
	}
}
