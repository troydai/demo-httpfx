package server

import (
	"net/http"

	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/fx"
)

type Param struct {
	fx.In

	Handlers []types.HttpHandler `group:"handlers"`
}

func NewHTTPServer(param Param) *http.Server {
	for _, h := range param.Handlers {
		http.HandleFunc(h.Path(), h.Handle)
	}

	return &http.Server{
		Addr: ":8080",
	}
}
