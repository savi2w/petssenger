package main

import (
	"github.com/weslenng/petssenger/services/user/grpc"
	"github.com/weslenng/petssenger/services/user/models"
	"github.com/weslenng/petssenger/services/user/redis"
)

func main() {
	db := models.InitDB()
	defer db.Close()
	defer redis.Client.Close()

	lis, err := grpc.UserServerListen()
	if err != nil {
		panic(err)
	}

	defer lis.Close()
}
