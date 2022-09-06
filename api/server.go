package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"github.com/go-redis/redis/v9"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/itsjoniur/currency/internal/utils"
	"github.com/itsjoniur/currency/internal/middlewares"
)

func StartAPI(logger *utils.Logger, storage *redis.Client, port string) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middlewares.CurrencyLogger(logger))
	r.Use(middlewares.RedisMiddleware(storage))
	r.Use(middlewares.RenderMiddleware(render.New()))
	r.Use(middleware.Recoverer)

	r.Get("/", RootHandler)
	r.Get("/price", CurrencyHandler)

	fmt.Printf("Server running on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Panicln(err)
	}
}
