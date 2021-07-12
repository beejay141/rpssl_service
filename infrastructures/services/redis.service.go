package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"rpssl_service/config"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

type RedisService struct {}

type IRedisService interface {
	Get(key string) *redis.StringCmd
	Set(key string, data interface{},exp int) error
	Remove(key string) (int64,error)
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: config.GetString("REDIS_HOST","localhost:6379"),
		Password: config.GetString("REDIS_PASSWORD",""),
	})
}




func (service RedisService) Get(key string) *redis.StringCmd {
	return rdb.Get(ctx, key)
}

// save data to redis, expiration is in seconds
func (service RedisService) Set(key string, data interface{},exp int) error {
	return rdb.Set(ctx, key, data, time.Duration(exp) * time.Second).Err()
}

func (service RedisService) Remove(key string) (int64,error) {
	return rdb.Del(ctx,key).Result()
}