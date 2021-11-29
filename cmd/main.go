package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("fail to provision a logger: %s", err)
	}


	logger.Info("starting a HTTP server ...")
	defer func() {
		logger.Info("stopping ...")
	}()
}
