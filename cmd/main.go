package main

import (
	"context"
	"net/http"

	"github.com/troydai/demo-httpfx/database"
	"github.com/troydai/demo-httpfx/logging"
	"github.com/troydai/demo-httpfx/routes/data"
	"github.com/troydai/demo-httpfx/routes/echo"
	"github.com/troydai/demo-httpfx/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			logging.NewLogger,
			database.NewDataSource,
			server.NewHTTPServer,
			data.NewDataHandler,
			echo.NewEchoHandler,
		),
		fx.Invoke(
			StartServer,
		),
	).Run()
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
