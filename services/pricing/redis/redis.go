package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/savi2w/petssenger/services/pricing/config"
)

var Client = redis.NewClient(&redis.Options{
	Addr: config.Default.RedisAddr,
})
