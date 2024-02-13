package main

import (
	"database/sql"
	"fmt"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/config"
)

func NewSqlCon(c config.Config) *sql.DB {
	con, err := sql.Open("postgres", c.DbUrl)
	if err != nil {
		panic(fmt.Errorf(
			"Unable to init sql connection using connection string %v. Error: %w",
			c.DbUrl,
			err,
		))
	}

	return con
}
