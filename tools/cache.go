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
	err      error
}

func getClient(ctx context.Context, server, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	cacheStatus := client.Ping(ctx)
	if err := cacheStatus.Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func NewCache(server, password string, db int) *Cache {
	ctx := context.Background()
	rdb, err := getClient(ctx, server, password, db)
	cache := &Cache{
		client:   rdb,
		server:   server,
		password: password,
		db:       db,
		err:      err,
	}
	return cache
}

func (cache *Cache) Get(key string) (string, error) {
	ctx := context.Background()
	if cache.err != nil {
		client, err := getClient(ctx, cache.server, cache.password, cache.db)
		if err != nil {
			return "", err
		}
		cache.client = client
	}
	val, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		err := fmt.Errorf("empty result for cache %s:%s", key, err.Error())
		return "", err
	}
	return val, nil
}

func (cache *Cache) SetWithTime(key, value string, expiredInSeconds int64, params ...int) error {
	ctx := context.Background()
	if cache.err != nil {
		client, err := getClient(ctx, cache.server, cache.password, cache.db)
		if err != nil {
			return err
		}
		cache.client = client
	}
	duration := time.Duration(expiredInSeconds * int64(time.Second))
	err := cache.client.Set(ctx, key, value, duration).Err()
	if err != nil {
		err := fmt.Errorf("Cache Set %s-%s:%s", key, value, err.Error())
		return err
	}
	return nil
}

func (cache *Cache) Set(key, value string, params ...int) error {
	ctx := context.Background()
	if cache.err != nil {
		client, err := getClient(ctx, cache.server, cache.password, cache.db)
		if err != nil {
			return err
		}
		cache.client = client
	}
	err := cache.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		err := fmt.Errorf("Cache Set %s-%s:%s", key, value, err.Error())
		return err
	}
	return nil
}
