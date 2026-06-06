// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"fmt"
	"speedsterApi/app/user/api/internal/config"
	handler2 "speedsterApi/app/user/api/internal/handler"
	"speedsterApi/app/user/api/internal/svc"
	"speedsterApi/common/casbin"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "etc/dev.yaml", "the config file")

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

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
	handler2.RegisterHandlers(server, ctx)
	handler2.RegisterDocRoute(server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
