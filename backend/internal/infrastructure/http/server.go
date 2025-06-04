package http

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/go-chi/chi"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetgrpc "github.com/Rintarooo/sample-connect-grpc/internal/component/greet/interface/grpc"
	"github.com/Rintarooo/sample-connect-grpc/internal/generated/grpc/greet/v1/greetv1connect"
	"github.com/Rintarooo/sample-connect-grpc/internal/infrastructure/middleware"
)

var server http.Server

func RunServer(logger *slog.Logger) error {
	mux := chi.NewRouter()

	greetHandler := greetgrpc.NewGreetingServiceHandler()

	path, handler := greetv1connect.NewGreetingServiceHandler(greetHandler)
	mux.Handle(path+"*", handler)

	reflector := grpcreflect.NewStaticReflector(
		"greet.v1.GreetingService",
	)
	mux.Mount(grpcreflect.NewHandlerV1(reflector))
	mux.Mount(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server = http.Server{
		Addr: ":8080",
		Handler: middleware.WithDefaultCORS()(
			h2c.NewHandler(mux, &http2.Server{}),
		),
	}

	return server.ListenAndServe()
}

func ShutdownServer(ctx context.Context) error {
	return server.Shutdown(ctx)
}
