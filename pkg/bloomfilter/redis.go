package bloomfilter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisStorage struct {
	client *redis.Client
	key    string
}

func NewRedisStorage(client *redis.Client, key string) *RedisStorage {
	return &RedisStorage{
		client: client,
		key:    key,
	}
}

func (r *RedisStorage) Store(locs []uint64) {
	ctx := context.Background()
	_, err := r.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, loc := range locs {
			log.Println(loc)
			pipe.SetBit(ctx, r.key, int64(loc), 1)
		}
		return nil
	})
	if err != nil {
		log.Println("redis storage bit err : ", err)
	}
}

func (r *RedisStorage) IsSet(loc uint64) bool {
	cmd := r.client.GetBit(context.Background(), r.key, int64(loc))
	return cmd.Val() == 1
}
