package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/itsjoniur/currency/api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/itsjoniur/currency/internal/utils"
)

const (
	port = "8000"
)

var (
	ctx = context.Background()
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

	logger := utils.NewLogger(logrus.New())

	api.StartAPI(logger, rdb, port)

}
