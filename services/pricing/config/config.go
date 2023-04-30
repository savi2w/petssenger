package config

import (
	"time"

	"github.com/go-pg/pg/v9"
)

type Config struct {
	Addr                    string
	DynamicFeesIncreaseRate float32
	DynamicFeesDecreaseRate float32
	DynamicFeesDecreaseTime time.Duration
	DynamicFeesMinimumValue float32
	PgConnOpts              pg.Options
	RedisAddr               string
	RedisExpTime            time.Duration
	WorkerRetryDelay        time.Duration
	WorkerRetryLimit        int
}

var Default = Config{
	Addr:                    "0.0.0.0:50051",
	DynamicFeesIncreaseRate: 0.1,
	DynamicFeesDecreaseRate: 0.1,
	DynamicFeesDecreaseTime: 5 * time.Minute,
	DynamicFeesMinimumValue: 1,
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
