package server

import (
	"context"
	"net/http"

	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/fx"
	"go.uber.org/zap"
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

func StartServer(lc fx.Lifecycle, server *http.Server, logger *zap.Logger) error {
	done := make(chan error)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				done <- server.ListenAndServe()
			}()
			logger.Info("starting a HTTP server. ctrl+c to exist.")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping server.")
			if err := server.Shutdown(ctx); err != nil {
				return err
			}
			<-done
			logger.Info("server terminated.")
			return nil
		},
	})

	return nil
}
