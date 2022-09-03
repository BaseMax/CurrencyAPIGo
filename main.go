package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

const (
	port = "8000"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("can not load .env file")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	status := rdb.Ping(ctx)
	ok, _ := status.Result()
	if ok != "PONG" {
		log.Panic("can not connect to redis")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(RedisMiddleware(rdb))
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
