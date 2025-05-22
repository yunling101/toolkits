package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type cache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedis(addr, password string, db int) *cache {
	return &cache{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       db,       // use default DB
			PoolSize: 10,
		}),
		ctx: context.Background(),
	}
}

func (c *cache) Ping() error {
	return c.client.Ping(c.ctx).Err()
}

func (c *cache) Set(key string, value interface{}) error {
	return c.client.Set(c.ctx, key, value, 0).Err()
}

func (c *cache) SetEx(key string, value interface{}, second int) error {
	return c.client.Set(c.ctx, key, value, time.Duration(second)*time.Second).Err()
}

func (c *cache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c *cache) Delete(key ...string) error {
	return c.client.Del(c.ctx, key...).Err()
}

func (c *cache) Close() error {
	return c.client.Close()
}
