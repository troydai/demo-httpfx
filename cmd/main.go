package main

import (
	"io"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("fail to provision a logger: %s", err)
	}

	echoHandler := func(w http.ResponseWriter, req *http.Request) {
		logger.Info("inbound", zap.String("path", req.URL.Path), zap.String("method", req.Method))

		body, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Error("failed to read request body", zap.Error(err))
			w.WriteHeader(500)
		}

		w.Write(body)
	}

	http.HandleFunc("/echo", echoHandler)

	logger.Info("starting a HTTP server ...")
	defer func() {
		logger.Info("stopping ...")
	}()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error("server exit with error", zap.Error(err))
	}
}
