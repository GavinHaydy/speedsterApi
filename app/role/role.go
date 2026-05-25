// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"fmt"
	"speedsterApi/common/casbin"

	"speedsterApi/app/role/internal/config"
	"speedsterApi/app/role/internal/handler"
	"speedsterApi/app/role/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/role-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//casbin
	svcCtx := svc.NewServiceContext(c)

	if err := casbin.Init(); err != nil {
		panic(err)
	}

	if err := casbin.LoadPolicy(svcCtx.DB); err != nil {
		panic(err)
	}
	// end

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	handler.RegisterDocRoute(server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
