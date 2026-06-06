// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AesSecretKey string
	DB           struct {
		DSN string // 对应 YAML 中的 DSN
	}
	Redis redis.RedisConf
	Auth  struct { // JWT 认证需要的密钥和过期时间配置
		Issuer        string
		Prefix        string
		AccessSecret  string
		AccessExpire  int
		RefreshSecret string
		RefreshExpire int
	}
	UserRpc zrpc.RpcClientConf
}
