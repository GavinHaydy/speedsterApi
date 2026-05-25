package handler

import (
	"fmt"
	"gateway/internal/svc"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/zeromicro/go-zero/rest"
)

type RouteConfig struct {
	Path   string // 匹配路径前缀
	Target string // 目标服务地址
}

// RegisterRoutes 注册所有的路由（包括文档聚合和业务接口转发）
func RegisterRoutes(engine *rest.Server, ctx *svc.ServiceContext) {
	// 修改点：将 ctx 传递给 RegisterDocRoutes

	var routes = []RouteConfig{
		{Path: "/user/", Target: ctx.Config.UserService.Target},
		{Path: "/role/", Target: ctx.Config.RoleService.Target},
		{Path: "/product/", Target: "http://product-service:8080"},
	}
	RegisterDocRoutes(engine, ctx)

	//业务路由转发逻辑保持不变...
	for _, route := range routes {
		path := route.Path + "/:path"
		path2 := route.Path + "/:path/:path"

		target, _ := url.Parse(route.Target)
		proxy := httputil.NewSingleHostReverseProxy(target)

		engine.AddRoutes([]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    path,
				Handler: proxyHandler(proxy),
			},
			{
				Method:  http.MethodGet,
				Path:    path2,
				Handler: proxyHandler(proxy),
			},
		})

		engine.AddRoutes([]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    path,
				Handler: proxyHandler(proxy),
			},
			{
				Method:  http.MethodPost,
				Path:    path2,
				Handler: proxyHandler(proxy)},
		})
		engine.AddRoutes([]rest.Route{
			{Method: http.MethodPut,
				Path:    path,
				Handler: proxyHandler(proxy)}, {
				Method:  http.MethodPut,
				Path:    path2,
				Handler: proxyHandler(proxy),
			},
		})
		engine.AddRoutes([]rest.Route{
			{Method: http.MethodDelete,
				Path:    path,
				Handler: proxyHandler(proxy)},
			{
				Method:  http.MethodDelete,
				Path:    path2,
				Handler: proxyHandler(proxy),
			},
		})
	}
}

// RegisterDocRoutes 专门用来注册 Scalar 文档聚合页面的路由
// 修改点：同样增加 ctx *svc.ServiceContext 参数
func RegisterDocRoutes(engine *rest.Server, ctx *svc.ServiceContext) {
	// 提供 Scalar 文档聚合页面
	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/docs",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Speedster API 文档</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
<div id="app"></div>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
	<script>
		Scalar.createApiReference('#app', {
			sources: [
    // API #1
    {
      title: 'User-Server',
      url: '/docs/user.json',
    },
	{
	title: 'Role-Server', 
	url: '/docs/role.json'
	},
]
})
	</script>
    
</body>
</html>
			`))
		},
	})

	// 代理各个微服务的 swagger.json 文件
	// 注意：这里暂时没用到 ctx，但为了保持函数签名一致先加上。
	// 以后如果你需要从配置文件（ctx.Config）中读取后端服务的真实地址，就可以在这里使用了。
	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/docs/user.json",
		//Handler: proxyTo("http://localhost:8888/docs/user.json"),
		Handler: proxyTo(fmt.Sprintf("%s/docs/user.json", ctx.Config.UserService.Target)),
	})
	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/docs/role.json",
		//Handler: proxyTo("http://order-service:8080/docs/swagger.json"),
		Handler: proxyTo(fmt.Sprintf("%s/docs/role.json", ctx.Config.RoleService.Target)),
	})
	engine.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/docs/product/swagger.json",
		Handler: proxyTo("http://product-service:8080/docs/swagger.json"),
	})
}

// proxyHandler 处理业务接口的反向代理
func proxyHandler(proxy *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

// proxyTo 处理 swagger.json 文档文件的反向代理
func proxyTo(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(target)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, resp.Body)
	}
}
