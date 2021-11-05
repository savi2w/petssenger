package main

import (
	"context"

	"github.com/omppye-lab/petssenger/services/pricing/grpc"
	"github.com/omppye-lab/petssenger/services/pricing/models"
	"github.com/omppye-lab/petssenger/services/pricing/redis"
	"github.com/omppye-lab/petssenger/services/pricing/worker"
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
