package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v9"
)

func StartAPI(storage *redis.Client, port string) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(RedisMiddleware(storage))
	r.Use(middleware.Recoverer)

	r.Get("/", RootHandler)
	r.Get("/price", CurrencyHandler)

	fmt.Printf("Server running on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Panicln(err)
	}
}

func RedisMiddleware(storage *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), 1, storage)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
