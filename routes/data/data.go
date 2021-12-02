package data

import (
	"io/ioutil"
	"net/http"

	"github.com/troydai/demo-httpfx/types"
	"go.uber.org/zap"
)

var _ types.HttpHandler = (*handler)(nil)

type handler struct {
	logger *zap.Logger
	source types.DataRetriver
}

func (h handler) Path() string {
	return "/data"
}

func (h handler) Handle(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("inbound", zap.String("path", req.URL.Path), zap.String("method", req.Method))

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "invalid body", 400)
		return
	}

	w.Write([]byte(h.source.Get(body)))
}

func NewDataHandler(logger *zap.Logger, source types.DataRetriver) types.HttpHandler {
	return &handler{
		logger: logger,
		source: source,
	}
}
