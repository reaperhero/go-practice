package redis

import (
	redis "gopkg.in/redis.v4"
)

var rds *redis.Client

func init() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
}