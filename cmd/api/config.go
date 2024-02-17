package main

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/config"
)

func NewConfig() config.Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	cfg := config.Config{}

	err = env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
