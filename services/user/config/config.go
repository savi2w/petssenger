package config

import (
	"time"

	"github.com/go-pg/pg/v9"
)

type Config struct {
	Addr         string
	HTTPPort     string
	PgConnOpts   pg.Options
	RedisAddr    string
	RedisExpTime time.Duration
}

var Default = Config{
	Addr:     "0.0.0.0:50051",
	HTTPPort: ":3002",
	PgConnOpts: pg.Options{
		Addr:     "user-postgres:5432",
		User:     "postgres",
		Password: "52b44f2327094ed59790a7506df7e1db",
		Database: "user",
	},
	RedisAddr:    "user-redis:6379",
	RedisExpTime: 1 * time.Minute,
}
