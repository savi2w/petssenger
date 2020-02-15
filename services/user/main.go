package main

import (
	"sync"

	"github.com/weslenng/petssenger/services/user/grpc"
	"github.com/weslenng/petssenger/services/user/http"
	"github.com/weslenng/petssenger/services/user/models"
	"github.com/weslenng/petssenger/services/user/redis"
)

func startGRPC() {
	lis, err := grpc.UserRPCListen()
	if err != nil {
		panic(err)
	}

	defer lis.Close()
}

func main() {
	db := models.InitDB()
	defer db.Close()
	defer redis.Client.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		startGRPC()
	}()

	go func() {
		defer wg.Done()
		http.UserHTTPListen()
	}()

	wg.Wait()
}
