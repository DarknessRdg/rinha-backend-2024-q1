package main

import (
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(NewRouter),
		fx.Provide(NewHTTPServer),
		fx.Provide(NewConfig),
		fx.Provide(NewSqlCon),
		fx.Provide(NewIAccountRepo),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
