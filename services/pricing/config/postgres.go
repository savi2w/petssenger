package config

import "github.com/go-pg/pg/v9"

// PostgresConnect to a docker postgres
func PostgresConnect() *pg.DB {
	conn := pg.Connect(&pg.Options{
		Addr:     "pricing-postgres:5432",
		User:     "postgres",
		Password: "04b76987bf0649519ef42136419d442e",
		Database: "pricing",
	})

	return conn
}
