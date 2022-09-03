package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
)

func LoadDataFromCache(key string) (map[string]string, error) {
	result := map[string]string{}

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(val), &result)
	return result, nil
}

func CacheData(key string, value any, ttl time.Duration) (bool, error) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false, err
	}

	exp := rdb.Expire(ctx, key, ttl*time.Second)
	ok, _ := exp.Result()
	if ok == false {
		rdb.Del(ctx, key)
		return ok, err
	}

	return true, nil
}
