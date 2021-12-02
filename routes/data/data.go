package data

import (
	"io/ioutil"
	"net/http"

	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/zap"
)

func NewDataHandler(logger *zap.Logger, source types.DataRetriver) types.HttpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info("inbound", zap.String("path", req.URL.Path), zap.String("method", req.Method))

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "invalid body", 400)
			return
		}

		w.Write([]byte(source.Get(body)))
	}
}
