package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"time"
)

var _cache *Cache

type Cache struct {
	ctx context.Context
	rdb *redis.Client
}

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("cache.addr"),
		Password: viper.GetString("cache.password"),
		DB:       viper.GetInt("cache.db"),
	})
	_cache = &Cache{
		ctx: context.Background(),
		rdb: rdb,
	}
}

func Instance() *Cache {
	return _cache
}

func (c *Cache) GetClient() *redis.Client {
	return c.rdb
}

func (c *Cache) Get(key string) string {
	log.Println("key : ", key)
	val, err := c.rdb.Get(c.ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Println("redis : ", err)
		panic(err)
	}
	return val
}

func (c *Cache) Set(key string, val string, exp time.Duration) {
	err := c.rdb.Set(c.ctx, key, val, exp).Err()
	if err != nil {
		panic(err)
	}
}
