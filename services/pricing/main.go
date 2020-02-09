package main

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"github.com/weslenng/petssenger/services/pricing/config"
	"github.com/weslenng/petssenger/services/pricing/server"
)

type context struct {
	pg    *pg.DB
	redis *redis.Client
}

func (ctx context) startServer() {
	lis, ser, err := server.PricingServerListen(ctx.pg, ctx.redis)
	if err != nil {
		panic(err)
	}

	defer lis.Close()
	defer ser.GracefulStop()
}

func main() {
	ctx := &context{
		pg:    config.PricingPgConnect(),
		redis: config.PricingRedisClient(),
	}

	ctx.startServer()
	defer ctx.pg.Close()
	defer ctx.redis.Close()
}
