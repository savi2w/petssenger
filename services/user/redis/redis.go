package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/omppye-lab/petssenger/services/user/config"
)

var Client = redis.NewClient(&redis.Options{
	Addr: config.Default.RedisAddr,
})
