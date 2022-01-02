package tools

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Cache struct {
	client   *redis.Client
	server   string
	password string
	db       int
}

func NewCache(server, password string, db int) (*Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	cache := &Cache{
		client:   rdb,
		server:   server,
		password: password,
		db:       db,
	}
	ctx := context.Background()
	cacheStatus := cache.client.Ping(ctx)
	if err := cacheStatus.Err(); err != nil {
		return nil, err
	}
	return cache, nil
}

func (cache *Cache) Get(key string) (string, error) {
	val, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		err := fmt.Errorf("empty result for cache %s:%s", key, err.Error())
		return "", err
	}
	return val, nil
}

func (cache *Cache) SetWithTime(key, value string, expiredInSeconds int64, params ...int) error {
	duration := time.Duration(expiredInSeconds)
	err := cache.client.Set(ctx, key, value, duration).Err()
	if err != nil {
		err := fmt.Errorf("Cache Set %s-%s:%s", key, value, err.Error())
		return err
	}
	return nil
}

func (cache *Cache) Set(key, value string, params ...int) error {
	err := cache.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		err := fmt.Errorf("Cache Set %s-%s:%s", key, value, err.Error())
		return err
	}
	return nil
}
