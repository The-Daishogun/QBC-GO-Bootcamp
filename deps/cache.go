package deps

import "github.com/redis/go-redis/v9"

type Caches struct {
	Calculations, Responses *redis.Client
}

func NewRedisClient(dbNumber int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       dbNumber,
	})
}
