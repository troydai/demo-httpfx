package echo

import (
	"io"
	"net/http"

	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var _ types.HttpHandler = (*handler)(nil)

type (
	handler struct {
		logger *zap.Logger
	}

	Result struct {
		fx.Out

		Handler types.HttpHandler `group:"handlers"`
	}
)

func (h handler) Path() string {
	return "/echo"
}

func (h handler) Handle(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("inbound", zap.String("path", req.URL.Path), zap.String("method", req.Method))

	body, err := io.ReadAll(req.Body)
	if err != nil {
		h.logger.Error("failed to read request body", zap.Error(err))
		w.WriteHeader(500)
	}

	w.Write(body)
}

func NewEchoHandler(logger *zap.Logger) Result {
	return Result{
		Handler: &handler{logger: logger},
	}
}
