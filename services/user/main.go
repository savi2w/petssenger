package main

import (
	"sync"

	"github.com/savi2w/petssenger/services/user/grpc"
	"github.com/savi2w/petssenger/services/user/http"
	"github.com/savi2w/petssenger/services/user/models"
	"github.com/savi2w/petssenger/services/user/redis"
)

func main() {
	db := models.InitDB()
	defer db.Close()
	defer redis.Client.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		lis, err := grpc.UserRPCListen()
		if err != nil {
			panic(err)
		}

		defer lis.Close()
	}()

	go func() {
		defer wg.Done()
		lis, err := http.UserHTTPListen()
		if err != nil {
			panic(err)
		}

		defer lis.Close()
	}()

	wg.Wait()
}
