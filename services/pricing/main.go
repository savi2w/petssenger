package main

import (
	"sync"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"github.com/weslenng/petssenger/services/pricing/config"
	"github.com/weslenng/petssenger/services/pricing/server"
	"github.com/weslenng/petssenger/services/pricing/worker"
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

func (ctx context) startWorker() {
	err := worker.MainQueue.Consumer().Start(nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := &context{
		pg:    config.PricingPgConnect(),
		redis: config.PricingRedisClient(),
	}

	defer ctx.pg.Close()
	defer ctx.redis.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ctx.startServer()
	}()

	go func() {
		defer wg.Done()
		ctx.startWorker()
	}()

	wg.Wait()
}
