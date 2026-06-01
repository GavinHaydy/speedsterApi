package middleware

import (
	"net/http"
	"speedsterApi/common/errno"
	"speedsterApi/common/response"
	"speedsterApi/common/utils"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisJwtMiddleware struct {
	RedisClient *redis.Redis
	Signature   string
}

func NewRedisJwtMiddleware(redisClient *redis.Redis, signature string) *RedisJwtMiddleware {
	return &RedisJwtMiddleware{
		RedisClient: redisClient,
		Signature:   signature,
	}
}

func (m *RedisJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//// 1. 从 Context 中获取 user_id（因为该中间件放在 @jwt 之后执行，此时 user_id 已经被注入）
		//userID := r.Context().Value("user_id")
		//logx.Infof("=========userID:%s===", userID)
		//
		//// 2. 拼接 Redis Key，去 Redis 中查询该 Token 是否还存在
		//tokenKey := fmt.Sprintf("%s%s", m.Prefix, userID)
		//exists, err := m.RedisClient.Exists(tokenKey)
		//
		//// 3. 如果 Redis 中不存在（返回 false 或报错），说明用户已退出或被强制下线
		//if err != nil || !exists {
		//	// 返回 401 未授权
		//	w.WriteHeader(http.StatusUnauthorized)
		//	w.Header().Set("Content-Type", "application/json")
		//	w.Write([]byte(`{"code": 401, "msg": "登录已失效，请重新登录"}`))
		//	return
		//}
		//
		//token, err := m.RedisClient.Get(tokenKey)
		//if err != nil {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		//
		//if token != r.Header.Get("Authorization") {
		//	w.Header().Set("Content-Type", "application/json")
		//	w.Write([]byte(`{"code": 401, "msg": "登录已失效，请重新登录"}`))
		//	return
		//}
		token := r.Header.Get("Authorization")
		_, _, tokenType, err := utils.ParseToken(token, m.Signature)
		if err != nil {
			//w.Header().Set("Content-Type", "application/json")
			//w.Write([]byte(`{"code": 401, "msg": "登录已失效，请重新登录"}`))
			response.ErrorWithCode(w, r, errno.ErrInvalidToken)
			return
		}
		if tokenType != "access" {
			//w.Header().Set("Content-Type", "application/json")
			//w.Write([]byte(`{"code": 401, "msg": "无效Token"}`))
			response.ErrorWithCode(w, r, errno.ErrTokenTypeFailed)
			return
		}
		// 4. Redis 校验通过，放行请求
		next(w, r)
	}
}
