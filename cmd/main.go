package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/troydai/demo-httpfx/database"
	"github.com/troydai/demo-httpfx/logging"
	"github.com/troydai/demo-httpfx/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := logging.NewLogger()
	if err != nil {
		log.Fatalf("fail to provision a logger: %s", err)
	}

	dataSource := database.NewDataSource()

	server := server.NewHTTPServer(logger, dataSource)
	Wait(StartServer(server, logger), logger)
}

func StartServer(server *http.Server, logger *zap.Logger) <-chan error {
	done := make(chan error)
	go func() {
		done <- server.ListenAndServe()
	}()

	logger.Info("starting a HTTP server. ctrl+c to exist.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		logger.Info("stopping server.")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	return done
}

func Wait(done <-chan error, logger *zap.Logger) {
	<-done
	logger.Info("server terminated.")
}
