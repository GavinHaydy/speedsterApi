// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	AesSecretKey string `json:",optional"`
	DB           struct {
		DSN string // 对应 YAML 中的 DSN
	}
	Redis redis.RedisConf
	JWT   struct {
		Secret string `json:",optional"`
		Issuer string `json:",optional"`
		Expire int    `json:",optional"`
	}
}
