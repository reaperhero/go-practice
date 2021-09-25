package redis

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"

	"github.com/go-redis/redis"
)

const (
	defaultExpireDuration = time.Minute * 15
)

// Cache ...
type Cache struct {
	redisClient *redis.Client
}

// NewCache ...
func NewCache(addr, pass string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pass,
		DB:           0,
		PoolSize:     300,
		MinIdleConns: 3,
		PoolTimeout:  time.Second * 15,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	})

	return &Cache{client}
}

// Set k-v in redis
// if expire <0, use defaultExpireDuration(15 minutes)
// if expire == 0, it will not expire
func (c Cache) Set(key string, value interface{}, expire time.Duration) error {
	if c.redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}
	cmd := c.redisClient.Do("SET", key, string(data))
	if cmd.Err() != nil {
		return fmt.Errorf("redis set error: %v", cmd.Err())
	}

	if expire < 0 {
		c.redisClient.Expire(key, defaultExpireDuration)
	} else if expire > 0 {
		c.redisClient.Expire(key, expire)
	}

	return nil
}

// Get ...
func (c Cache) Get(key string, value interface{}) error {
	if c.redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	result := c.redisClient.Get(key)
	if result.Err() != nil {
		return fmt.Errorf("redis get error: %v", result.Err())
	}

	err := json.Unmarshal([]byte(result.Val()), value)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %v", err)
	}

	return nil
}

// Delete ...
func (c Cache) Delete(key string) error {
	if c.redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	result := c.redisClient.Del(key)
	if result.Err() != nil {
		return fmt.Errorf("redis delete error: %v", result.Err())
	}
	return nil
}

func (c Cache) Lpush(key string, value interface{}) (lLen int64, err error) {
	if c.redisClient == nil {
		return 0, fmt.Errorf("redis client is nil")
	}
	data, err := json.Marshal(value)
	if err != nil {
		return 0, fmt.Errorf("json marshal error: %v", err)
	}
	cmd := c.redisClient.LPush(key, data)
	if cmd.Err() != nil {
		return 0, fmt.Errorf("redis lpush error: %v", cmd.Err())
	}

	return cmd.Val(), nil
}

func (c Cache) LTRIM(key string, start, stop int64) (err error) {
	if c.redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}
	cmd := c.redisClient.LTrim(key, start, stop)
	if cmd.Err() != nil {
		return fmt.Errorf("redis LTrim error: %v", cmd.Err())
	}

	return nil
}
func (c Cache) LRem(key string, count int64, value interface{}) (err error) {
	if c.redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal error: %v", err)
	}
	log.Info(data)
	cmd := c.redisClient.LRem(key, count, data)
	if cmd.Err() != nil {
		return fmt.Errorf("redis LRem error: %v", cmd.Err())
	}

	return nil
}

func (c Cache) LRange(key string, start, stop int64) (result []string, err error) {
	if c.redisClient == nil {
		return nil, fmt.Errorf("redis client is nil")
	}
	cmd := c.redisClient.LRange(key, start, stop)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("redis LRem error: %v", cmd.Err())
	}

	return cmd.Val(), nil
}

func (c Cache) HMSet(key string, fields map[string]interface{}, expire time.Duration) error {
	if c.redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	cmd := c.redisClient.HMSet(key, fields)
	if cmd.Err() != nil {
		return fmt.Errorf("redis LRem error: %v", cmd.Err())
	}

	if expire < 0 {
		c.redisClient.Expire(key, defaultExpireDuration)
	} else if expire > 0 {
		c.redisClient.Expire(key, expire)
	}
	return nil
}
