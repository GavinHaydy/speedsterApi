package handler

import (
	"embed"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

//go:embed docs/*.json
var docFS embed.FS

func RegisterDocRoute(server *rest.Server) {
	server.AddRoutes([]rest.Route{
		{
			Method: http.MethodGet,
			Path:   "/docs/role.json",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				data, _ := docFS.ReadFile("docs/role.json")
				w.Header().Set("Content-Type", "application/json")
				if _, err := w.Write(data); err != nil {
					return
				}
			},
		},
	})
}
