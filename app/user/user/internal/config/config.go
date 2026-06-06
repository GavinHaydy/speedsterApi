package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DSN string // 对应 YAML 中的 DSN
	}
	CacheRedis redis.RedisConf
	CacheAuth  struct { // JWT 认证需要的密钥和过期时间配置
		Issuer        string
		Prefix        string
		AccessSecret  string
		AccessExpire  int
		RefreshSecret string
		RefreshExpire int
	}
}
