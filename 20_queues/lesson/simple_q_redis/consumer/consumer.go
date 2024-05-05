package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
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
		arrayVal, err := redisCli.BRPop(context.Background(), 0, "my_list_key").Result() // 0 timeout - wait for messages forever
		if err != nil {
			log.Fatalf("error reading: %s", err)
		}

		log.Printf("got: %v", arrayVal[1])
	}
}
