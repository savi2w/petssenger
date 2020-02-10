package main

import (
	"context"
	"sync"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"github.com/weslenng/petssenger/services/pricing/config"
	"github.com/weslenng/petssenger/services/pricing/server"
	"github.com/weslenng/petssenger/services/pricing/worker"
)

type environment struct {
	pg    *pg.DB
	redis *redis.Client
}

func (env environment) startServer() {
	lis, ser, err := server.PricingServerListen(env.pg, env.redis)
	if err != nil {
		panic(err)
	}

	defer lis.Close()
	defer ser.GracefulStop()
}

func (env environment) startWorker() {
	err := worker.MainQueue.Consumer().Start(context.Background())
	if err != nil {
		panic(err)
	}
}

func main() {
	env := &environment{
		pg:    config.PricingPgConnect(),
		redis: config.PricingRedisClient(),
	}

	defer env.pg.Close()
	defer env.redis.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		env.startServer()
	}()

	go func() {
		defer wg.Done()
		env.startWorker()
	}()

	defer worker.MainQueue.Close()
	wg.Wait()
}
