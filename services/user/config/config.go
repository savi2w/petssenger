package config

import (
	"time"

	"github.com/go-pg/pg/v9"
)

// Config is a structure for a common configuration of a user service
type Config struct {
	Addr         string
	PgConnOpts   pg.Options
	RedisAddr    string
	RedisExpTime time.Duration
}

// Default contains a common configuration for a user service
var Default = Config{
	Addr: "0.0.0.0:50051",
	PgConnOpts: pg.Options{
		Addr:     "user-postgres:5432",
		User:     "postgres",
		Password: "04b76987bf0649519ef42136419d442e",
		Database: "user",
	},
	RedisAddr:    "user-redis:6379",
	RedisExpTime: 1 * time.Minute,
}
