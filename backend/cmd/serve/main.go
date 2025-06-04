package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rintarooo/sample-connect-grpc/internal/infrastructure/http"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("app", "sample-connect-grpc")

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-quitCh
		logger.Info("shutting down server ...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		if err := http.ShutdownServer(shutdownCtx); err != nil {
			logger.Error("failed to shutdown server", "error", err)
			os.Exit(1)
			return
		}
	}()

	logger.Info("server started ...")
	if err := http.RunServer(logger); err != nil {
		logger.Error("failed to run server", "error", err)
		os.Exit(1)
	}
}
