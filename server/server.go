package server

import (
	"net/http"

	"github.com/troydai/demo-httpfx/routes/data"
	"github.com/troydai/demo-httpfx/routes/echo"
	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/zap"
)

func NewHTTPServer(logger *zap.Logger, dataSource types.DataRetriver) *http.Server {
	handlers := []types.HttpHandler{
		echo.NewEchoHandler(logger),
		data.NewDataHandler(logger, dataSource),
	}

	for _, h := range handlers {
		http.HandleFunc(h.Path(), h.Handle)
	}

	return &http.Server{
		Addr: ":8080",
	}
}
