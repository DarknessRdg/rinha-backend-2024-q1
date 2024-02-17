package main

import (
	"database/sql"
	"fmt"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/config"
	// We need to import postgresql driver, so the `sql.DB` can load
	// the actual db implementation
	_ "github.com/lib/pq"
)

func NewSqlCon(c config.Config) *sql.DB {
	con, err := sql.Open("postgres", c.DbUrl)
	if err != nil {
		panic(fmt.Errorf(
			"unable to init sql connection using connection string %v. Error: %w",
			c.DbUrl,
			err,
		))
	}

	err = con.Ping()
	if err != nil {
		panic(err)
	}

	return con
}
