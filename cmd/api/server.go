package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewHTTPServer(router chi.Router, lc fx.Lifecycle) *http.Server {
	server := &http.Server{Addr: "127.0.0.1:3000", Handler: router}
	log := slog.With("address", server.Addr)

	log.Info("instantiating HTTP Server")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listen, err := net.Listen("tcp", server.Addr)
			if err != nil {
				log.Error("errors listening")
				return err
			}

			log.Info("starting HTTP server")
			go server.Serve(listen)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("shutting down server")
			return server.Shutdown(ctx)
		},
	})
	return server
}
