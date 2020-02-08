package postgres

import (
	"github.com/go-pg/pg/v9"
)

// Connect to a docker database
func Connect() *pg.DB {
	conn := pg.Connect(&pg.Options{
		Addr:     "pricing-database:5432",
		User:     "postgres",
		Password: "04b76987bf0649519ef42136419d442e",
		Database: "pricing",
	})

	return conn
}
