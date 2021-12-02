package routes

import (
	"github.com/troydai/demo-httpfx/routes/data"
	"github.com/troydai/demo-httpfx/routes/echo"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		data.NewDataHandler,
		echo.NewEchoHandler,
	),
)
