package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"time"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		DB:       0, // use default DB
		Protocol: 3, // specify 2 for RESP 2 or 3 for RESP 3
	})
	err := redisCli.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}

	defer redisCli.Close()

	for {
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		val := time.Now().Unix()
		redisCli.LPush(context.Background(), "my_list_key", val)
		log.Printf("published: %d", val)
	}
}
