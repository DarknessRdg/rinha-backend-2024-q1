package main

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/config"

func NewConfig() config.Config {
	return config.Config{
		DbUrl: "",
		Port:  3000,
	}
}
