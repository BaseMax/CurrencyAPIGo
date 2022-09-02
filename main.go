package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = "8000"
)

func main() {

	status := rdb.Ping(ctx)
	ok, _ := status.Result()
	if ok != "PONG" {
		log.Panic("can not connect to redis")
	}

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/price", CurrencyHandler)

	fmt.Printf("Server running on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Panicln(err)
	}
}
