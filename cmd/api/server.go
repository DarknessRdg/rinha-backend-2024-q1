package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewHTTPServer(router chi.Router, lc fx.Lifecycle) *http.Server {
	server := &http.Server{Addr: ":3000", Handler: router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listen, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", server.Addr)
			go server.Serve(listen)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}
