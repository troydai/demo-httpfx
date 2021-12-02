package main

import (
	"github.com/troydai/demo-httpfx/database"
	"github.com/troydai/demo-httpfx/logging"
	"github.com/troydai/demo-httpfx/routes"
	"github.com/troydai/demo-httpfx/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.Module,
		routes.Module,
		fx.Provide(
			logging.NewLogger,
			database.NewDataSource,
		),
	).Run()
}
