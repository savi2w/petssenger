package config

import (
	"time"

	"github.com/go-pg/pg/v9"
)

// Config is a structure for a common configuration of a service
type Config struct {
	Addr                    string
	DynamicFeesMinimumValue float32
	DynamicFeesIncreaseRate float32
	DynamicFeesDecreaseRate float32
	DynamicFeesDecreaseTime time.Duration
	PgConnOpts              pg.Options
	RedisAddr               string
	RedisExpTime            time.Duration
	WorkerRetryDelay        time.Duration
	WorkerRetryLimit        int
}

// Default contains a common configuration for a service
var Default = Config{
	Addr:                    "0.0.0.0:50051",
	DynamicFeesMinimumValue: 1,
	DynamicFeesDecreaseRate: 0.1,
	DynamicFeesIncreaseRate: 0.1,
	DynamicFeesDecreaseTime: 5 * time.Second,
	PgConnOpts: pg.Options{
		Addr:     "pricing-postgres:5432",
		User:     "postgres",
		Password: "04b76987bf0649519ef42136419d442e",
		Database: "pricing",
	},
	RedisAddr:        "pricing-redis:6379",
	RedisExpTime:     1 * time.Minute,
	WorkerRetryDelay: 3 * time.Second,
	WorkerRetryLimit: 3,
}
