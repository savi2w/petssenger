package config

import (
	"github.com/go-pg/pg/v9"
)

// Config is a structure for a common configuration of a service
type Config struct {
	Addr       string
	PgConnOpts pg.Options
	RedisAddr  string
}

// Default contains a common configuration for a service
var Default = Config{
	Addr: "0.0.0.0:50051",
	PgConnOpts: pg.Options{
		Addr:     "pricing-postgres:5432",
		User:     "postgres",
		Password: "04b76987bf0649519ef42136419d442e",
		Database: "pricing",
	},
	RedisAddr: "pricing-redis:6379",
}
