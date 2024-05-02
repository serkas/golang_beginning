package main

import (
	"context"
	"log"
	"proj/lessons/19_redis/lesson/cached/internal/app"
)

func main() {
	conf := app.Config{
		ServerAddress: ":8081",
		DB:            "root:root@tcp(localhost:23306)/items_db",
	}
	a, err := app.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting server at %s", conf.ServerAddress)
	if err := a.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

}
