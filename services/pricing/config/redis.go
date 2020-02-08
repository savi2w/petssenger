package config

import "github.com/go-redis/redis/v7"

// RedisClient connect to a docker redis
func RedisClient() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     "pricing-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return cli
}
