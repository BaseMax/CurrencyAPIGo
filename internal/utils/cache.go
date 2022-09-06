package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
)

func LoadDataFromCache(storage *redis.Client, key string) (map[string]string, error) {
	result := map[string]string{}

	val, err := storage.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(val), &result)
	return result, nil
}

func CacheData(storage *redis.Client, key string, value any, ttl time.Duration) (bool, error) {
	err := storage.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return false, err
	}

	exp := storage.Expire(context.Background(), key, ttl*time.Second)
	ok, _ := exp.Result()
	if ok == false {
		storage.Del(context.Background(), key)
		return false, err
	}

	return true, nil
}
