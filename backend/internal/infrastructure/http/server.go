package http

import (
	"context"
	"log/slog"
	"net/http"
)

var server http.Server

func RunServer(logger *slog.Logger) error {
	server = http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	return server.ListenAndServe()
}

func ShutdownServer(ctx context.Context) error {
	return server.Shutdown(ctx)
}
