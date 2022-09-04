package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

func LoadDataFromCache(storage *redis.Client, key string) (map[string]string, error) {
	result := map[string]string{}

	val, err := storage.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(val), &result)
	return result, nil
}

func CacheData(storage *redis.Client, key string, value any, ttl time.Duration) (bool, error) {
	err := storage.Set(ctx, key, value, 0).Err()
	if err != nil {
		return false, err
	}

	exp := storage.Expire(ctx, key, ttl*time.Second)
	ok, _ := exp.Result()
	if ok == false {
		storage.Del(ctx, key)
		return ok, err
	}

	return true, nil
}
