package main

import (
	"log/slog"

	"github.com/DarknessRdg/rinha-backend-2024-q1/cmd/api/endpoints"
	"github.com/go-chi/chi/v5"
)

func NewRouter(
	transactionEndpoint *endpoints.TransactionEndpoint,
) chi.Router {
	slog.Info("instantiating routers")

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		transactionEndpoint.Router(r)
	})

	return r
}
