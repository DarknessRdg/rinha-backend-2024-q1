package main

import (
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New().Run()

	fx.New(
		fx.Provide(NewRouter),
		fx.Provide(NewHTTPServer),
		fx.Provide(NewConfig),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
