package utils

import "github.com/go-redis/redis"

var (
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GetYmlProperties().Redis.Url,
		Password: GetYmlProperties().Redis.Password,
		DB:       GetYmlProperties().Redis.Db,
	})
}
