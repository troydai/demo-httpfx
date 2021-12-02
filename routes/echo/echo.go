package echo

import (
	"io"
	"net/http"

	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/zap"
)

func NewEchoHandler(logger *zap.Logger) types.HttpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info("inbound", zap.String("path", req.URL.Path), zap.String("method", req.Method))

		body, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Error("failed to read request body", zap.Error(err))
			w.WriteHeader(500)
		}

		w.Write(body)
	}
}
