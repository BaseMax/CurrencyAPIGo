package middlewares

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v9"
)

// RedisMiddleware put redis client into context
func RedisMiddleware(storage *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), 1, storage)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
