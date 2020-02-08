package config

import "github.com/go-redis/redis/v7"

// PricingRedisClient is a helper function to connect in a Redis CLI
func PricingRedisClient() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr: "pricing-redis:6379",
	})

	return cli
}
