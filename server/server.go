package server

import (
	"net/http"

	"github.com/troydai/demo-httpfx/routes/data"
	"github.com/troydai/demo-httpfx/routes/echo"
	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/zap"
)

func NewHTTPServer(logger *zap.Logger, dataSource types.DataRetriver) *http.Server {
	http.HandleFunc("/echo", echo.NewEchoHandler(logger))
	http.HandleFunc("/data", data.NewDataHandler(logger, dataSource))

	return &http.Server{
		Addr: ":8080",
	}
}
