package main

import (
	"net/http"

	"github.com/DarknessRdg/rinha-backend-2024-q1/cmd/api/endpoints"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewRouter),
		fx.Provide(NewHTTPServer),
		fx.Provide(NewConfig),
		fx.Provide(NewSqlCon),
		fx.Provide(NewIAccountRepo),
		fx.Provide(NewITransactionRepo),
		fx.Provide(NewITransactionService),
		fx.Provide(endpoints.NewTransactionEndpoint),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
