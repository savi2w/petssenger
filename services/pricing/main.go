package main

import (
	"context"

	"github.com/weslenng/petssenger/services/pricing/grpc"
	"github.com/weslenng/petssenger/services/pricing/models"
	"github.com/weslenng/petssenger/services/pricing/redis"
	"github.com/weslenng/petssenger/services/pricing/worker"
)

func main() {
	db := models.InitDB()
	defer db.Close()
	defer redis.Client.Close()

	if err := worker.MainQueue.Consumer().Start(context.Background()); err != nil {
		panic(err)
	}

	defer worker.MainQueue.Close()
	lis, err := grpc.PricingRPCListen()
	if err != nil {
		panic(err)
	}

	defer lis.Close()
}
