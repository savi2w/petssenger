package config

import "github.com/go-pg/pg/v9"

// PricingPgConnect is a helper function to connect in a PostgreSQL database
func PricingPgConnect() *pg.DB {
	conn := pg.Connect(&pg.Options{
		Addr:     "pricing-postgres:5432",
		User:     "postgres",
		Password: "04b76987bf0649519ef42136419d442e",
		Database: "pricing",
	})

	return conn
}
